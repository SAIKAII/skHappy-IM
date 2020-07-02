package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Group struct {
	GroupId      uint64     `gorm:"group_id,omitempty"`
	GroupName    string     `gorm:"group_name"`
	CreateUser   string     `gorm:"create_user"`
	Owner        string     `gorm:"owner"`
	Announcement string     `gorm:"announcement,omitempty"`
	UserNum      int        `gorm:"user_num,omitempty"`
	IsDeleted    int8       `gorm:"is_deleted,omitempty"`
	CreatedAt    *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"updated_at,omitempty"`
}

type GroupUser struct {
	Id        uint64     `gorm:"id,omitempty"`
	GroupId   uint64     `gorm:"group_id"`
	Username  string     `gorm:"username"`
	IsDeleted int8       `gorm:"is_deleted,omitempty"`
	CreatedAt *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty"`
}

type GroupDao struct {
	DB *gorm.DB
}

type GroupUserDao struct {
	DB *gorm.DB
}

// InsertOne 这个接口一定要以事务方式执行
func (g *GroupDao) InsertOne(group *Group) (uint64, error) {
	if err := g.DB.Create(group).Error; err != nil {
		return 0, err
	}
	var id []uint64
	g.DB.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)

	return id[0], nil
}

func (g *GroupDao) GetOne(groupId uint64) (*Group, error) {
	group := &Group{}
	if err := g.DB.First(group, "group_id = ?", groupId).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (g *GroupDao) DeleteOne(groupId uint64) error {
	if err := g.DB.Table("groups").Where("group_id = ?", groupId).Update("is_deleted", 1).Error; err != nil {
		return err
	}

	return nil
}

func (g *GroupDao) UpdateNum(groupId uint64, num int) error {
	if err := g.DB.Table("groups").Where("group_id = ?", groupId).Update("user_num", num).Error; err != nil {
		return err
	}

	return nil
}

func (g *GroupDao) UserNum(groupId uint64) (int, error) {
	var num []int
	if err := g.DB.Table("groups").Where("group_id = ?", groupId).Pluck("user_num", &num).Error; err != nil {
		return 0, err
	} else if num == nil {
		return 0, DAO_ERROR_RECORD_NOT_FOUND
	}

	return num[0], nil
}

func (gu *GroupUserDao) InsertOne(groupUser *GroupUser) error {
	if err := gu.DB.Create(groupUser).Error; err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return DAO_ERROR_DUPLICATE_RECORD
		}
		return err
	}

	return nil
}

func (gu *GroupUserDao) GetOne(groupId uint64, username string) (*GroupUser, error) {
	groupUser := &GroupUser{}
	if err := gu.DB.Where("group_id = ? and username = ?", groupId, username).First(groupUser).Error; err != nil {
		return nil, err
	}

	return groupUser, nil
}

func (gu *GroupUserDao) UpdateOne(groupUser *GroupUser) error {
	if err := gu.DB.Table("group_users").
		Where("group_id = ? and username = ?", groupUser.GroupId, groupUser.Username).
		Update("is_deleted", groupUser.IsDeleted).Error; err != nil {
		return err
	}

	return nil
}

func (gu *GroupUserDao) GetAll(groupId uint64) ([]*GroupUser, error) {
	users := make([]*GroupUser, 0)
	if err := gu.DB.Find(&users, "group_id = ? and is_deleted = 0", groupId).Error; err != nil {
		return nil, err
	}

	return users, nil
}
