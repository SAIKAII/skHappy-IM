package apis

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/jwt"
	"github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

var _ pb.ConnServiceServer = &ConnServer{}

type ConnServer struct {
}

func StartConnRPCServer(addr string) {
	connListen, err := net.Listen("tcp", addr)
	defer connListen.Close()
	if err != nil {
		panic(err)
	}

	connServer := grpc.NewServer(grpc.UnaryInterceptor(connInterceptor))
	pb.RegisterConnServiceServer(connServer, &ConnServer{})
	err = connServer.Serve(connListen)
	if err != nil {
		panic(err)
	}
}

// connInterceptor JWT认证
func connInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "从context中获取数据失败")
	}

	// 进行JWT获取与认证
	var jwtString string
	if v, ok := md["jwt"]; ok {
		jwtString = v[0]
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "获取JWT失败")
	}

	_, err = jwt.VerifyJWT(jwtString)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return handler(ctx, req)
}

func (c *ConnServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	err := services.IMessageService.SendToUser(ctx, req)
	if err != nil {
		if err == dao.DAO_ERROR_RECORD_NOT_FOUND {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DeliverMessageResp{}, nil
}
