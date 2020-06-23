package service

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/encrypto"
	"github.com/SAIKAII/skHappy-IM/pkg/uuid"
	"github.com/SAIKAII/skHappy-IM/services"
	"time"
)

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) error {
	// 验证帐号是否已存在
	_, err := a.GetAccount(dto.Username)
	if err != nil && err != dao.DAO_ERROR_RECORD_NOT_FOUND {
		return err
	}

	user := &dao.User{
		Username: dto.Username,
		Nickname: dto.Nickname,
		Password: dto.Password,
		Avatar:   dto.Avatar,
		Sex:      dto.Sex,
		Birthday: dto.Birthday,
		PhoneNum: dto.PhoneNum,
	}
	// TODO 两个数据库操作做成事务
	err = a.create(user)
	if err != nil {
		return err
	}

	err = a.initMsgRecv(dto.Username)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountService) GetAccount(username string) (*services.AccountReturnDTO, error) {
	db := base.Database()
	accountDao := dao.UserDao{DB: db}
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

	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	users, err := accountDao.GetAll(allUsers)
	if err != nil {
		return nil, err
	}

	rUsers := make([]*services.AccountReturnDTO, len(users))
	for i, u := range users {
		rUsers[i] = &services.AccountReturnDTO{
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Birthday: time.Unix(0, 0),
		}
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
	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	err = accountDao.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountService) generateSalt() (string, error) {
	return uuid.NewUUID()
}

func (a *accountService) initMsgRecv(username string) error {
	msgRecv := &dao.MsgRecv{
		Username:  username,
		LastSeqId: 0,
	}

	db := base.Database()
	msgRecvDao := dao.MsgRecvDao{DB: db}
	err := msgRecvDao.InsertOne(msgRecv)
	if err != nil {
		return err
	}

	return nil
}
