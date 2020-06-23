package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id        uint64     `gorm:"id,omitempty"`
	Username  string     `gorm:"username,uni"`
	Nickname  string     `gorm:"nickname"`
	Password  string     `gorm:"password"`
	Salt      string     `gorm:"salt"`
	Avatar    string     `gorm:"avatar,omitempty"`
	Sex       int8       `gorm:"sex,omitempty"` // 0-unknown 1-male 2-female
	Birthday  time.Time  `gorm:"birthday,omitempty"`
	PhoneNum  string     `gorm:"phone_num,omitempty"`
	IsDeleted int8       `gorm:"is_deleted,omitempty"`
	CreatedAt *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty"`
}

type UserDao struct {
	DB *gorm.DB
}

func (a *UserDao) GetOne(username string) (*User, error) {
	userInfo := &User{}
	if err := a.DB.First(&userInfo, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, DAO_ERROR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	return userInfo, nil
}

func (a *UserDao) Insert(u *User) error {
	if err := a.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (a *UserDao) GetAll(allUsers []string) ([]*User, error) {
	users := make([]*User, 0)
	if err := a.DB.Where("username IN (?)", allUsers).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (a *UserDao) UpdateProfile(u *User) error {
	if err := a.DB.Model(&u).Updates(map[string]interface{}{
		"nickname":  u.Nickname,
		"avatar":    u.Avatar,
		"sex":       u.Sex,
		"birthday":  u.Birthday,
		"phone_num": u.PhoneNum,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (a *UserDao) UpdatePassword(u *User) error {
	if err := a.DB.Model(&u).Updates(map[string]interface{}{
		"salt":     u.Salt,
		"password": u.Password,
	}).Error; err != nil {
		return err
	}

	return nil
}
