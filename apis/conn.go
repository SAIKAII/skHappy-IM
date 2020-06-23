package apis

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"google.golang.org/grpc"
	"net"
)

var _ pb.ConnServiceServer = &ConnServer{}

type ConnServer struct {
}

func StartConnRPCServer(addr string) {
	connListen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	connServer := grpc.NewServer()
	pb.RegisterConnServiceServer(connServer, &ConnServer{})
	err = connServer.Serve(connListen)
	if err != nil {
		panic(err)
	}
}

func (c *ConnServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	err := services.IMessageService.SendToOne(req)
	if err != nil {
		return nil, err
	}

	return &pb.DeliverMessageResp{}, nil
}
