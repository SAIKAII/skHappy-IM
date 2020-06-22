package service

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/encrypto"
	"github.com/SAIKAII/skHappy-IM/pkg/uuid"
	"github.com/SAIKAII/skHappy-IM/services"
)

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) error {
	// 验证帐号是否已存在
	_, err := a.GetAccount(dto.Username)
	if err != nil {
		return err
	}

	user := &dao.User{
		Username: dto.Username,
		Password: dto.Password,
		Sex:      dto.Sex,
		Birthday: dto.Birthday,
		PhoneNum: dto.PhoneNum,
	}
	err = a.create(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountService) GetAccount(username string) (*services.AccountReturnDTO, error) {
	accountDao := dao.UserDao{}
	accountDao.DB = base.Database()
	user, err := accountDao.GetOne(username)
	if err != nil {
		return nil, err
	}

	return &services.AccountReturnDTO{
		Username: user.Username,
		Sex:      user.Sex,
		Birthday: user.Birthday,
		PhoneNum: user.PhoneNum,
	}, nil
}

func (a *accountService) GetAccounts(username string) ([]*services.AccountReturnDTO, error) {
	allUsers, err := services.IRelationshipService.GetAllFriends(username)
	if err != nil {
		return nil, err
	}

	accountDao := dao.UserDao{}
	users, err := accountDao.GetAll(allUsers)
	if err != nil {
		return nil, err
	}

	rUsers := make([]*services.AccountReturnDTO, len(users))
	for i, u := range users {
		rUsers[i].Nickname = u.Nickname
		rUsers[i].Avatar = u.Avatar
	}

	return rUsers, nil
}

func (a *accountService) UpdateProfile(dto services.AccountUpdateDTO) error {
	panic("")
}

func (a *accountService) create(user *dao.User) error {
	var err error
	if user.Salt, err = a.generateSalt(); err != nil {
		return err
	}

	user.Password = encrypto.EnCryptoPassword(user.Password, user.Salt)
	accountDao := dao.UserDao{}
	accountDao.DB = base.Database()
	err = accountDao.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountService) generateSalt() (string, error) {
	return uuid.NewUUID()
}
