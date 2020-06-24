package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MsgRecv struct {
	Id        uint64     `gorm:"id,omitempty"`
	Username  string     `gorm:"username"`
	LastSeqId uint64     `gorm:"last_seq_id"`
	CreatedAt *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty"`
}

type MsgRecvDao struct {
	DB *gorm.DB
}

func (mr *MsgRecvDao) GetOne(username string) (*MsgRecv, error) {
	msgRecv := &MsgRecv{}
	if err := mr.DB.First(&msgRecv, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, DAO_ERROR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return msgRecv, nil
}

func (mr *MsgRecvDao) InsertOne(msgRecv *MsgRecv) error {
	if err := mr.DB.Create(msgRecv).Error; err != nil {
		return err
	}

	return nil
}

func (mr *MsgRecvDao) UpdateLastSeqId(username string, lastSeqId uint64) error {
	if err := mr.DB.Table("msg_recvs").Where("username = ?", username).Update("last_seq_id", lastSeqId).Error; err != nil {
		return err
	}
	return nil
}
