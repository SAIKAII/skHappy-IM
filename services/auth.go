package services

import "errors"

var (
	USER_NOT_FOUND = errors.New("没有该用户")
	AUTH_FAILURE   = errors.New("认证失败")
)

var IAuthService AuthService

type AuthService interface {
	SignInAuth(username, pwd string) error
}
