package services

import (
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	pb "github.com/SAIKAII/skHappy-IM/protocols"
)

var IGroupService GroupService

type GroupService interface {
	CreateGroup(*pb.CreateGroupReq) (uint64, error)
	DeleteGroup(uint64) error
	AddGroupMember(*pb.AddGroupMemberReq) error
	DelGroupMember(*pb.DelGroupMemberReq) error
	ListGroupMember(groupId uint64) ([]*dao.GroupUser, error)
	IsMember(groupId uint64, username string) (bool, error)
}
