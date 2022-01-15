分析的源代码版本： v1.7-agent

源代码分析文章会做成一个系列，以下是思路。

1、agent和server的通信流程剖析

2、agent和driver的通信流程剖析

3、agent和插件plugin的通信

4、plugin功能插件都有哪些？找一个典型的进行分析（待定）

。。。。。。

源代码如下所示：

# 一、proto文件

agent的proto文件路径：Elkeid/agent/proto/grpc.proto

server的proto文件路径：Elkeid/server/agent_center/grpctrans/proto

agent端核心结构如下所示：

```protobuf
message PackagedData {
  repeated EncodedRecord records = 1;  	#已编码的记录列表
  string agent_id = 2;
  repeated string intranet_ipv4 = 3;	#内网ipv4地址
  repeated string extranet_ipv4 = 4;	#外网ipv4地址
  repeated string intranet_ipv6 = 5;	#内网ipv6地址
  repeated string extranet_ipv6 = 6;	#外网ipv6地址
  string hostname = 7;					#主机名
  string version = 8;					#版本
  string product = 9;					#产品名
}
```



# 二、agent端源码分析

## 2、1 调用堆栈

agent端属于client端，想创建client，必然要调用grpc.pb.go文件中的函数NewTransferClient

```go
func NewTransferClient(cc *grpc.ClientConn) TransferClient {
   return &transferClient{cc}
}
```

![](Elkeid_code%20analysis.assets/image-20220115223239201.png)

调用者为agent/transport目录下，transfer.go文件的startTransfer函数

```go
func startTransfer(ctx context.Context, wg *sync.WaitGroup) {
   defer wg.Done()
   retries := 0
   subWg := &sync.WaitGroup{}
   defer subWg.Wait()
   for {
      //判断连接是否存在
      conn := connection.GetConnection(ctx)
      // 获取不到连接，则记录重试次数，当重试次数大于5时，提示无可用连接
      if conn == nil {
         if retries > 5 {
            zap.S().Error("transfer will shutdown because of no avaliable connections")
            return
         }
         //等待5秒中，再去获取下一个连接connection
         zap.S().Warnf("wait to get next connection for 5 seconds,current retry times:%v", retries)
         select {
         case <-ctx.Done():
            return
         case <-time.After(time.Second * 5):
            retries++
            continue
         }
      }
      zap.S().Infof("get connection successfully:idc %v,region %v,netmode %v", connection.IDC, connection.Region, connection.NetMode.Load().(string))
      retries = 0
      var client proto.Transfer_TransferClient
      subCtx, cancel := context.WithCancel(ctx)
      //创建grpc客户端对象，并调用Transfer
      client, err := proto.NewTransferClient(conn).Transfer(subCtx, grpc.UseCompressor("snappy"))
      if err == nil {
         subWg.Add(2)
         //数据发送逻辑
         go handleSend(subCtx, subWg, client)
         go func() {
            //数据接收逻辑
            handleReceive(subCtx, subWg, client)
            // 收到错误后取消服务
            cancel()
         }()
         subWg.Wait()
      } else {
         zap.S().Error(err)
      }
      cancel()
      zap.S().Info("transfer has been canceled,wait next try to transfer for 5 seconds")
      #传输被取消，等待5秒后尝试下一次
      select {
      case <-ctx.Done():
         return
      case <-time.After(time.Second * 5):
      }
   }
}
```

看这里如果忽略错误处理代码的话，我们关心的函数仅仅有三个：

1、connection.GetConnection()

2、处理数据发送逻辑的函数 handleSend()

3、处理数据接收逻辑的函数 handleReceive()

## 2、2 核心函数

### 2、2、1 connection.GetConnection函数





### 2、2、2 handleSend()函数

handleSend()函数中使用的全局变量如下：

```go
var (
   Mu                = &sync.Mutex{}		//多个client同时读写Buf，保证其协程安全的互斥锁
   Buf               = [8192]interface{}{}	//数据发送缓冲区
   Offset            = 0					//发送缓冲区的数据偏移量
   ErrBufferOverflow = errors.New("buffer overflow") //超过发送缓冲区的错误
   hook              func(interface{}) interface{}
   RecordPool        = sync.Pool{			//EncodedRecord类型的对象池
      New: func() interface{} {
         return &proto.EncodedRecord{
            Data: make([]byte, 0, 1024*2),
         }
      },
   }
)
```

**Mu：**多个client同时读写Buf，保证其协程安全的互斥锁

**Buf：**数据发送缓冲区

**Offset：**发送缓冲区的数据偏移量

**ErrBufferOverflow：**超过发送缓冲区容量时，会产生错误

**RecordPool：**EncodedRecord类型的对象池

此处RecordPool采用的对象池sync.Pool，对于频繁创建和销毁的对象使用对象池技术，能大大的提升程序性能。



函数中创建了一个0.1秒的定时器，即每隔0.1秒会触发一次数据发送事件，核心逻辑如下：

```go
case <-ticker.C:
   {
      //多个客户端同时操作同一个发送缓冲区core.Offset，所以需加锁
      core.Mu.Lock()
      //发送缓冲区有数据
      if core.Offset != 0 {
         zap.S().Debugf("will send %v recs", core.Offset)
         //创建EncodedRecord切片，用来存储EncodedRecord
         nbuf := make([]*proto.EncodedRecord, 0, core.Offset)
         for _, v := range core.Buf[:core.Offset] {
            //判断元素类型
            //proto.EncodedRecord类型的数据，直接添加到切片中
            //proto.Record类型的数据，转换为proto.EncodedRecord类型的数据再添加到切片中
            switch t := v.(type) {
            case *proto.EncodedRecord:
               nbuf = append(nbuf, t)
            case *proto.Record:
               data, _ := t.Data.Marshal()
               rec := core.RecordPool.Get().(*proto.EncodedRecord)
               rec.DataType = t.DataType
               rec.Timestamp = t.Timestamp
               rec.Data = data
               nbuf = append(nbuf, rec)
            }
         }
         //填充PackagedData结构体并发送出去
         err := client.Send(&proto.PackagedData{
            Records:      nbuf,
            AgentId:      agent.ID,
            IntranetIpv4: host.PrivateIPv4.Load().([]string),
            IntranetIpv6: host.PrivateIPv6.Load().([]string),
            ExtranetIpv4: host.PublicIPv4.Load().([]string),
            ExtranetIpv6: host.PublicIPv6.Load().([]string),
            Hostname:     host.Name.Load().(string),
            Version:      agent.Version,
            Product:      agent.Product,
         })
         //数据发送结束后，将对象归还给对象池core.RecordPool
         for _, v := range nbuf {
            v.Data = v.Data[:0]
            core.RecordPool.Put(v)
         }
         //发送成功则记录发送次数
         if err == nil {
            atomic.AddUint64(&txCnt, uint64(core.Offset))
            core.Offset = 0
         } else {
            core.Mu.Unlock()
            return
         }
      }
      core.Mu.Unlock()
   }
```

1、多个客户端同时操作同一个发送缓冲区core.Offset，所以需加锁Mu；

2、当发送缓冲区有数据时，创建EncodedRecord切片，用来存储EncodedRecord；

3、遍历发送缓冲区中数据并判断类型：

​		proto.EncodedRecord类型的数据，直接添加到切片中；

​		proto.Record类型的数据，转换为proto.EncodedRecord类型的数据再添加到切片中；

4、填充PackagedData结构体（如agentId，内网ipv4，nbuf记录结果等）并通过grpc接口发送出去

5、 数据发送结束后，将对象归还给对象池core.RecordPool

6、发送成功后，释放锁Mu，并进行统计计数。

Elkeid采用多个grpc连接复用一个发送缓冲区core.Buf，Offset用于指示发送缓冲区中的偏移量，当缓冲区的偏移量大于0时才会触发实际的数据发送操作。



### 2、2、3 handleReceive()函数

```go
func handleReceive(ctx context.Context, wg *sync.WaitGroup, client proto.Transfer_TransferClient) {
   defer wg.Done()
   defer zap.S().Info("receive handler will exit")
   zap.S().Info("receive handler running")
   for {
      cmd, err := client.Recv()
      if err != nil {
         zap.S().Error(err)
         return
      }
      zap.S().Info("received command")
      atomic.AddUint64(&rxCnt, 1) //统计接收数
      if cmd.Task != nil {
         // 给agent的任务
         if cmd.Task.ObjectName == agent.Product {
            switch cmd.Task.DataType {
            case 1060: //1060表示关闭agent
               zap.S().Info("will shutdown agent")
               agent.Cancel()
               zap.S().Info("shutdown agent successfully")
               return
            }

         } else {
            // 给插件的任务
            // 根据插件名称来获取插件对象，成功则向插件发送任务
            // 插件对象为空，则打印错误日志
            plg, ok := plugin.Get(cmd.Task.ObjectName)
            if ok {
               err = plg.SendTask(*cmd.Task)
               if err != nil {
                  plg.Error("send task to plugin failed: ", err)
               }
            } else {
               zap.S().Error("can't find plugin: ", cmd.Task.ObjectName)
            }
         }
         continue
      }
      // 处理配置变更
      cfgs := map[string]*proto.Config{}
      for _, config := range cmd.Configs {
         cfgs[config.Name] = config
      }
      // 配置的版本和agent的版本不同，则更新agent
      if cfg, ok := cfgs[agent.Product]; ok && cfg.Version != agent.Version {
         zap.S().Infof("agent will update:current version %v -> expected version %v", agent.Version, cfg.Version)
         err := agent.Update(*cfg)
         if err == nil {
            zap.S().Info("update successfully")
            agent.Cancel()
            return
         } else {
            zap.S().Error("update failed:", err)
         }
      }
      delete(cfgs, agent.Product)
      // 升级agent成功后同步plugin
      err = plugin.Sync(cfgs)
      if err != nil {
         zap.S().Error(err)
      }
      continue
   }
}
```

流程如下：

1、调用client.Recv()从server端接收消息

2、判断任务cmd.Task是否为nil，大体分为2种情况

3、给agent的任务，1060表示关闭agent

4、给插件的任务，根据插件名称来获取插件对象，插件对象有效则向插件发送任务

5、处理配置变更，当下发配置的版本和本地的agent版本不同时，更新agent并同步plugin插件信息。



服务端核心数据结构如下：



所以在不同平台下，我们还是要不断的突破自己。

然后聚焦在server和agent的创建函数上：

主要讲Agent和Server之间的通信过程