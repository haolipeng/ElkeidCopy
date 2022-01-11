package grpc_handler

import (
	"ElkeidCopy/server/agent_center/grpctrans/pool"
	pb "ElkeidCopy/server/agent_center/grpctrans/proto"
)

type TransferHandler struct{}

var (
	//GlobalGRPCPool is the global grpc connection manager
	GlobalGRPCPool *pool.GRPCPool
)

func InitGlobalGRPCPool() {

}

func (h *TransferHandler) Transfer(server pb.Transfer_TransferServer) error {
	//TODO:
	panic("implement me")
}
