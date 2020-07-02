package service

import (
	"github.com/SAIKAII/skHappy-IM/services"
	"sync"
)

var _ services.AccountService = &accountService{}
var _ services.RelationshipService = &relationshipService{}
var _ services.AuthService = &authService{}
var _ services.MessageService = &messageService{}
var _ services.GroupService = &groupService{}
var once sync.Once

func init() {
	once.Do(func() {
		services.IAccountService = &accountService{}
		services.IRelationshipService = &relationshipService{}
		services.IAuthService = &authService{}
		services.IMessageService = &messageService{}
		services.IGroupService = &groupService{}
	})
}
