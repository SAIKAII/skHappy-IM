package service

import (
	"errors"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/encrypto"
	"github.com/SAIKAII/skHappy-IM/pkg/uuid"
	"github.com/SAIKAII/skHappy-IM/services"
	"github.com/jinzhu/gorm"
	"time"
)

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) error {
	// 验证帐号是否已存在
	_, err := a.GetAccount(dto.Username)
	if err != nil && err != dao.DAO_ERROR_RECORD_NOT_FOUND {
		base.Logger.Errorln(err)
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

	db := base.Database()
	err = db.Transaction(func(tx *gorm.DB) error {
		err = a.create(user, tx)
		if err != nil {
			base.Logger.Errorln(err)
			return err
		}

		err = a.initMsgRecv(dto.Username, tx)
		if err != nil {
			base.Logger.Errorln(err)
			return err
		}

		return nil
	})
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (a *accountService) GetAccount(username string) (*services.AccountReturnDTO, error) {
	user, err := a.getUser(username)
	if err != nil {
		base.Logger.Errorln(err)
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
		base.Logger.Errorln(err)
		return nil, err
	}

	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	users, err := accountDao.GetAll(allUsers)
	if err != nil {
		base.Logger.Errorln(err)
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
	db := base.Database()
	userDao := dao.UserDao{DB: db}
	user, err := userDao.GetOne(dto.Username)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	user.Nickname = dto.Nickname
	user.Avatar = dto.Avatar
	user.Birthday = dto.Birthday
	user.PhoneNum = dto.PhoneNum
	user.Sex = dto.Sex
	err = userDao.UpdateProfile(user)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (a *accountService) ChangePassword(dto services.ChangePasswordDTO) error {
	user, err := a.getUser(dto.Username)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	res := encrypto.EnCryptoPassword(dto.PrePassword, user.Salt)
	if res != user.Password {
		err := errors.New("原密码不正确")
		base.Logger.Errorln(err)
		return err
	}

	if user.Salt, err = a.generateSalt(); err != nil {
		base.Logger.Errorln(err)
		return err
	}

	user.Password = encrypto.EnCryptoPassword(dto.Password, user.Salt)
	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	err = accountDao.UpdatePassword(user)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (a *accountService) getUser(username string) (*dao.User, error) {
	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	return accountDao.GetOne(username)
}

func (a *accountService) create(user *dao.User, db *gorm.DB) error {
	var err error
	if user.Salt, err = a.generateSalt(); err != nil {
		base.Logger.Errorln(err)
		return err
	}

	user.Password = encrypto.EnCryptoPassword(user.Password, user.Salt)
	accountDao := dao.UserDao{DB: db}
	err = accountDao.Insert(user)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}

func (a *accountService) generateSalt() (string, error) {
	return uuid.NewUUID()
}

func (a *accountService) initMsgRecv(username string, db *gorm.DB) error {
	msgRecv := &dao.MsgRecv{
		Username:  username,
		LastSeqId: 0,
	}

	msgRecvDao := dao.MsgRecvDao{DB: db}
	err := msgRecvDao.InsertOne(msgRecv)
	if err != nil {
		base.Logger.Errorln(err)
		return err
	}

	return nil
}
