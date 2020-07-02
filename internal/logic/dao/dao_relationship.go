package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Relationship struct {
	Id        int64      `gorm:"id,omitempty"`
	UserA     string     `gorm:"user_a"`
	UserB     string     `gorm:"user_b"`
	IsDeleted int8       `gorm:"is_deleted,omitempty"`
	CreatedAt *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty"`
}

// Greater 比较UserA是否应该排在UserB前
func (r *Relationship) Greater() bool {
	return r.UserA >= r.UserB
}

type RelationShipDao struct {
	DB *gorm.DB
}

// Insert 插入两人关系记录
func (r *RelationShipDao) Insert(rel *Relationship) error {
	if err := r.DB.Create(rel).Error; err != nil {
		return err
	}

	return nil
}

// Update 更新记录
func (r *RelationShipDao) Update(rel *Relationship) error {
	if err := r.DB.Table("relationships").Where("user_a = ? and user_b = ?",
		rel.UserA, rel.UserB).Update("is_deleted", rel.IsDeleted).Error; err != nil {
		return err
	}
	return nil
}

// GetAll 获取指定用户的所有好友信息
func (r *RelationShipDao) GetAll(username string) ([]*Relationship, error) {
	rels := make([]*Relationship, 0)
	if err := r.DB.Find(&rels, "user_a = ? or user_b = ?", username, username).Error; err != nil {
		return nil, err
	}

	return rels, nil
}

// Delete 删除两人关系
func (r *RelationShipDao) Delete(userA, userB string) error {
	if err := r.DB.Where("user_a = ? and user_b = ?", userA, userB).Update("is_deleted", 1).Error; err != nil {
		return err
	}

	return nil
}

// GetOne 获取指定两人的关系记录，若没有指定记录不会引发错误，但是返回的对象为nil
func (r *RelationShipDao) GetOne(userA, userB string) (*Relationship, error) {
	rel := &Relationship{}
	if err := r.DB.First(rel, "user_a = ? and user_b = ?", userA, userB).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, DAO_ERROR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return rel, nil
}
