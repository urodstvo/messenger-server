package message

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Delete(ctx context.Context, id uint32) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Message{}).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
