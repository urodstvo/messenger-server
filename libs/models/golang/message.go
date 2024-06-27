package models

import (
	"time"
)

type MessageStatus string

const (
	MessageCreated MessageStatus = "created"
	MessageReaded  MessageStatus = "readed"
)

type Message struct {
	ID        uint32        `gorm:"primary_key;autoIncrement" json:"id"`
	ChatID    uint32        `json:"chat_id"`
	AuthorID  uint32        `json:"author_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
	Text      string        `json:"text"`
	Status    MessageStatus `gorm:"default:created" json:"status"`

	Chat   Chat `gorm:"foreignkey:ChatID;references:ID" json:"chat"`
	Author User `gorm:"foreignkey:AuthorID;references:ID" json:"author"`
}

func (m *Message) TableName() string {
	return "messages"
}
