package grpctrans

import (
	"ElkeidCopy/server/agent_center/grpctrans/grpc_handler"
	pb "ElkeidCopy/server/agent_center/grpctrans/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

const (
	// Maximum connection idle time
	defaultMaxConnIdle = 20 * time.Minute

	//If the connection is idle during pingtime,
	//the server takes the initiative to ping http_client
	defaultPingTime = 10 * time.Minute

	//Same as above, the timeout period of server waiting for ack when pinging client
	defaultPingAckTimeout = 5 * time.Second

	maxMsgSize = 1024 * 1024 * 10 // grpc maximum message size:10M
	addr       = ":8080"
)

func Run() {
	//TODO:初始化连接池

	//启动server
	runServer()
}

func runServer() {
	//0.服务器的选项
	//keepalive server parameters
	kasp := keepalive.ServerParameters{
		MaxConnectionIdle: defaultMaxConnIdle,
		Time:              defaultPingTime,
		Timeout:           defaultPingAckTimeout,
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(kasp),
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.MaxRecvMsgSize(maxMsgSize),
	}

	//1、新建grpc服务器端
	grpcServer := grpc.NewServer(opts...)

	//2、调用pb接口注册对象
	pb.RegisterTransferServer(grpcServer, &grpc_handler.TransferHandler{})

	//反射注册
	reflection.Register(grpcServer)

	//3、tcp侦听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("net.Listen error: %s\n", err)
		return
	}

	fmt.Printf("runServer is running")

	//4、grpc server提供服务，需传net.Listener类型参数
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("grpcServer.Serve error: %s\n", err)
		return
	}
}
