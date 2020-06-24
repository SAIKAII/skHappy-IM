package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Message struct {
	Id           uint64     `gorm:"id,omitempty"`
	SenderType   int8       `gorm:"sender_type"`
	Sender       string     `gorm:"sender"`
	ReceiverType int8       `gorm:"receiver_type"`
	Receiver     string     `gorm:"receiver"`
	Type         int8       `gorm:"type"`
	Content      string     `gorm:"content"`
	SeqId        uint64     `gorm:"seq_id"`
	SendTime     *time.Time `gorm:"sender_time"`
	CreatedAt    *time.Time `gorm:"create_time,omitempty"`
	UpdatedAt    *time.Time `gorm:"update_time,omitempty"`
}

type MessageDao struct {
	DB *gorm.DB
}

func (m *MessageDao) GetOneBySeqId(seqId uint64) (*Message, error) {
	msg := &Message{}
	if err := m.DB.First(&msg, "seq_id = ?", seqId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, DAO_ERROR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	return msg, nil
}

func (m *MessageDao) InsertOne(msg *Message) error {
	if err := m.DB.Create(msg).Error; err != nil {
		return err
	}
	return nil
}
