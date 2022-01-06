package connection

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync/atomic"
	"time"
)

var (
	conn atomic.Value //*grpc.ClientConn
)

func GetConnection(ctx context.Context) *grpc.ClientConn {
	//1.conn是否存在连接
	c, ok := conn.Load().(*grpc.ClientConn)
	if ok {
		//TODO:待完善的代码，是否添加重试机制??
		//判断clientConn的状态
		switch c.GetState() {
		case connectivity.Ready: //ready
			return c
		case connectivity.Idle: //Idle
			return c
		case connectivity.Connecting: //connecting
			c.Close()
		case connectivity.TransientFailure: //TransientFailure
			c.Close()
		case connectivity.Shutdown: //shutdown
		}
	}

	//2.创建timeout context，默认三秒
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel() //在函数退出或超时触发，则调用cancel函数

	//3.调用grpc.DialContext
	target := "127.0.0.1:8088"
	c, err := grpc.DialContext(timeoutCtx, target)

	//4.将获得的clientConn保存到原子变量中
	if err == nil {
		conn.Store(c)
		return c
	}

	return nil
}
