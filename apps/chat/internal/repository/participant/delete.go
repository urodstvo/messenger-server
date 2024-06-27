package participant

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Delete(ctx context.Context, chatId, userId uint32) error {
	if err := r.db.WithContext(ctx).Where("chat_id = ?", chatId).Where("user_id = ?", userId).Delete(&models.Participant{}).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
