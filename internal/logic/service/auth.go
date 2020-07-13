package service

import (
	"errors"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/SAIKAII/skHappy-IM/internal/logic/dao"
	"github.com/SAIKAII/skHappy-IM/pkg/encrypto"
	"github.com/SAIKAII/skHappy-IM/services"
)

type authService struct {
}

func (a *authService) SignInAuth(username, password string) error {
	db := base.Database()
	accountDao := dao.UserDao{DB: db}
	user, err := accountDao.GetOne(username)
	if err != nil {
		err := errors.New("该用户未注册")
		base.Logger.Errorln(err)
		return err
	}

	pwdEncrypted := encrypto.EnCryptoPassword(password, user.Salt)
	if pwdEncrypted != user.Password {
		base.Logger.Errorln(services.AUTH_FAILURE)
		return services.AUTH_FAILURE
	}

	return nil
}
