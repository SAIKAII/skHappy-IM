package services

import (
	"time"
)

const (
	KEY_MEMBER_ONLINE = "online"
)

var IAccountService AccountService

type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) error
	GetAccount(username string) (*AccountReturnDTO, error)
	GetAccounts(username string) ([]*AccountReturnDTO, error)
	UpdateProfile(dto AccountUpdateDTO) error
}

type AccountCreatedDTO struct {
	Username string
	Nickname string
	Password string
	Avatar   string
	Sex      int8
	Birthday time.Time
	PhoneNum string
}

type AccountUpdateDTO struct {
	Nickname string
	Password string
	Avatar   string
	Sex      int8
	Birthday time.Time
	PhoneNum string
}

type AccountReturnDTO struct {
	Username string
	Nickname string
	Avatar   string
	Sex      int8
	Birthday time.Time
	PhoneNum string
}
