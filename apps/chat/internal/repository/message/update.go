package message

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Update(ctx context.Context, message *models.Message, columns ...string) error {
	if err := r.db.WithContext(ctx).Model(&message).Select(columns).Omit("id").Updates(message).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
