package chat

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Update(ctx context.Context, chat *models.Chat, columns ...string) error {
	if err := r.db.WithContext(ctx).Model(&chat).Select(columns).Omit("id").Updates(chat).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
