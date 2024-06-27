package models

import (
	"time"
)

type ChatRole string

const (
	ChatRoleOwner  ChatRole = "owner"
	ChatRoleAdmin  ChatRole = "admin"
	ChatRoleMember ChatRole = "member"
)

type Chat struct {
	ID        uint32     `gorm:"primary_key;autoIncrement" json:"id"`
	Name      string     `json:"name"`
	Tag       string     `gorm:"unique" json:"tag"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsPublic  bool       `json:"is_public"`

	Participants []Participant `gorm:"foreignkey:ChatID" json:"participants"`
	Messages     []Message     `gorm:"foreignkey:ChatID" json:"messages"`
}

func (Chat) TableName() string {
	return "chats"
}

type Participant struct {
	ChatID uint32   `json:"chat_id"`
	UserID uint32   `json:"user_id"`
	Role   ChatRole `gorm:"default:member" json:"role"`

	Chat Chat `gorm:"foreignkey:ChatID;references:ID" json:"chat"`
	User User `gorm:"foreignkey:UserID;references:ID" json:"user"`
}

func (Participant) TableName() string {
	return "participants"
}
