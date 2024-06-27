package message

import (
	def "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	"gorm.io/gorm"
)

var _ def.MessageRepository = (*repository)(nil)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
