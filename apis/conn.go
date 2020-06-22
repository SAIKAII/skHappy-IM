package apis

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
)

var _ pb.ConnServiceServer = &ConnServer{}

type ConnServer struct {
}

func (c ConnServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	err := services.IMessageService.SendToOne(req)
	if err != nil {
		return nil, err
	}

	return &pb.DeliverMessageResp{}, nil
}
