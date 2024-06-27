package message

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Create(ctx context.Context, message *models.Message) (uint32, error) {
	if err := r.db.WithContext(ctx).Create(&message).Error; err != nil {
		return 0, errors.NewRepositoryError(err)
	}

	return message.ID, nil
}
