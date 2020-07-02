package apis

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/jwt"
	"github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

var _ pb.CliInterfaceServiceServer = &CliInterfaceServer{}

type CliInterfaceServer struct {
}

func StartCliRPCServer(addr string) {
	cliListen, err := net.Listen("tcp", addr)
	defer cliListen.Close()
	if err != nil {
		panic(err)
	}

	efp := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}
	sp := keepalive.ServerParameters{
		MaxConnectionIdle:     60 * time.Second,
		MaxConnectionAge:      1 * time.Hour,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  10 * time.Second,
		Timeout:               1 * time.Second,
	}
	cliServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(efp),
		grpc.KeepaliveParams(sp),
		grpc.UnaryInterceptor(cliInterceptor))
	pb.RegisterCliInterfaceServiceServer(cliServer, &CliInterfaceServer{})
	err = cliServer.Serve(cliListen)
	if err != nil {
		panic(err)
	}
}

// cliInterceptor JWT认证
func cliInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.FullMethod != "/pb.CliInterfaceService/Register" {
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

		_, err := jwt.VerifyJWT(jwtString)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "jwt", jwtString)
	}

	return handler(ctx, req)
}

func (cf *CliInterfaceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	err := services.IAccountService.CreateAccount(services.AccountCreatedDTO{
		Username: req.User.Username,
		Nickname: req.User.Nickname,
		Password: req.User.Password,
		Avatar:   req.User.AvatarUrl,
		Sex:      int8(req.User.Sex),
		Birthday: time.Unix(req.User.Birthday, 0),
		PhoneNum: req.User.PhoneNum,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.RegisterResp{}, nil
}

func (cf *CliInterfaceServer) AddFriend(ctx context.Context, req *pb.AddFriendReq) (*pb.AddFriendResp, error) {
	err := services.IRelationshipService.CreateRelationship(req.UserId, req.FriendId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.AddFriendResp{}, nil
}

func (cf *CliInterfaceServer) GetFriend(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	rdto, err := services.IAccountService.GetAccount(req.Username)
	if err != nil {
		if err == dao.DAO_ERROR_RECORD_NOT_FOUND {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetUserResp{
		User: &pb.User{
			Username:  rdto.Username,
			Nickname:  rdto.Nickname,
			AvatarUrl: rdto.Avatar,
			Sex:       int32(rdto.Sex),
			Birthday:  rdto.Birthday.Unix(),
			PhoneNum:  rdto.PhoneNum,
		},
	}, nil
}

func (cf *CliInterfaceServer) ListFriends(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	rdtos, err := services.IAccountService.GetAccounts(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	users := make([]*pb.User, len(rdtos))
	for i, v := range rdtos {
		users[i] = &pb.User{
			Username:  v.Username,
			Nickname:  v.Nickname,
			AvatarUrl: v.Avatar,
			Sex:       int32(v.Sex),
			Birthday:  v.Birthday.Unix(),
			PhoneNum:  v.PhoneNum,
		}
	}

	return &pb.ListUsersResp{Users: users}, nil
}

func (cf *CliInterfaceServer) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	err := services.IMessageService.Send(ctx, req)
	if err != nil {
		if err == dao.DAO_ERROR_RECORD_NOT_FOUND {
			return nil, status.Errorf(codes.FailedPrecondition, "用户不在该群组内")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SendMessageResp{}, nil
}

func (cf *CliInterfaceServer) DelFriend(ctx context.Context, req *pb.DelFriendReq) (*pb.DelFriendResp, error) {
	err := services.IRelationshipService.DeleteRelationship(req.Username, req.FriendName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DelFriendResp{}, nil
}

func (cf *CliInterfaceServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileResp, error) {
	err := services.IAccountService.UpdateProfile(services.AccountUpdateDTO{
		Username: req.User.Username,
		Nickname: req.User.Nickname,
		Password: req.User.Password,
		Avatar:   req.User.AvatarUrl,
		Sex:      int8(req.User.Sex),
		Birthday: time.Unix(req.User.Birthday, 0),
		PhoneNum: req.User.PhoneNum,
	})
	if err != nil {
		if err == dao.DAO_ERROR_RECORD_NOT_FOUND {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UpdateProfileResp{}, nil
}

func (cf *CliInterfaceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.ChangePasswordResp, error) {
	err := services.IAccountService.ChangePassword(services.ChangePasswordDTO{
		Username:    req.Username,
		PrePassword: req.PrePassword,
		Password:    req.Password,
	})
	if err != nil {
		if err == dao.DAO_ERROR_RECORD_NOT_FOUND {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ChangePasswordResp{}, nil
}

func (cf *CliInterfaceServer) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	groupId, err := services.IGroupService.CreateGroup(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateGroupResp{GroupId: groupId}, nil
}

func (cf *CliInterfaceServer) DeleteGroup(ctx context.Context, req *pb.DisbandGroupReq) (*pb.DisbandGroupResp, error) {
	err := services.IGroupService.DeleteGroup(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DisbandGroupResp{}, nil
}

func (cf *CliInterfaceServer) AddGroupMember(ctx context.Context, req *pb.AddGroupMemberReq) (*pb.AddGroupMemberResp, error) {
	err := services.IGroupService.AddGroupMember(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddGroupMemberResp{}, nil
}

func (cf *CliInterfaceServer) DelGroupMember(ctx context.Context, req *pb.DelGroupMemberReq) (*pb.DelGroupMemberResp, error) {
	err := services.IGroupService.DelGroupMember(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DelGroupMemberResp{}, nil
}
