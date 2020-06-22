package apis

import (
	"context"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/protocols"
	"github.com/SAIKAII/skHappy-IM/services"
	"time"
)

var _ pb.CliInterfaceServiceServer = &CliInterfaceServer{}

type CliInterfaceServer struct {
}

func (cf *CliInterfaceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	services.IAccountService.CreateAccount(services.AccountCreatedDTO{
		Username: req.User.Username,
		Nickname: req.User.Nickname,
		Password: req.User.Password,
		Avatar:   req.User.AvatarUrl,
		Sex:      int8(req.User.Sex),
		Birthday: time.Unix(req.User.Birthday, 0),
		PhoneNum: req.User.PhoneNum,
	})
	return &pb.RegisterResp{}, nil
}

func (cf *CliInterfaceServer) AddFriend(ctx context.Context, req *pb.AddFriendReq) (*pb.AddFriendResp, error) {
	err := services.IRelationshipService.CreateRelationship(req.UserId, req.FriendId)
	if err != nil {
		return nil, err
	}

	return &pb.AddFriendResp{}, nil
}

func (cf *CliInterfaceServer) GetFriend(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	rdto, err := services.IAccountService.GetAccount(req.Username)
	if err != nil {
		return nil, err
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
		return nil, err
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
	rpcCli := base.NewRPCCli()
	rpcConn, err := rpcCli.Dialer(base.USER_ADDR, req.Item.ReceiverName)
	if err != nil {
		return nil, err
	}

	_, err = pb.NewConnServiceClient(rpcConn).DeliverMessage(ctx, &pb.DeliverMessageReq{
		Item: req.Item,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResp{}, nil
}

func (cf *CliInterfaceServer) DelFriend(ctx context.Context, req *pb.DelFriendReq) (*pb.DelFriendResp, error) {
	err := services.IRelationshipService.DeleteRelationship(req.Username, req.FriendName)
	if err != nil {
		return nil, err
	}

	return &pb.DelFriendResp{}, nil
}

func (cf *CliInterfaceServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileResp, error) {
	err := services.IAccountService.UpdateProfile(services.AccountUpdateDTO{
		Nickname: req.User.Nickname,
		Password: req.User.Password,
		Avatar:   req.User.AvatarUrl,
		Sex:      int8(req.User.Sex),
		Birthday: time.Unix(req.User.Birthday, 0),
		PhoneNum: req.User.PhoneNum,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResp{}, nil
}
