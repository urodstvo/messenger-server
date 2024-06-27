package message

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (r *repository) GetById(ctx context.Context, id uint32) (*models.Message, error) {
	var message models.Message
	if err := r.db.WithContext(ctx).Preload(clause.Associations).Where("id = ?", id).First(&message).Error; err != nil {
		return nil, errors.NewRepositoryError(err)
	}

	return &message, nil
}
