package chat

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (r *repository) GetById(ctx context.Context, id uint32) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).Preload(clause.Associations).Where("id = ?", id).First(&chat).Error; err != nil {
		return nil, errors.NewRepositoryError(err)
	}

	return &chat, nil
}
