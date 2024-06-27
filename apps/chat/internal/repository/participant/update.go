package participant

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Update(ctx context.Context, participant *models.Participant) error {
	if err := r.db.WithContext(ctx).Model(&models.Participant{}).Where("user_id = ?", participant.UserID).Where("chat_id = ?", participant.ChatID).Update("role", participant.Role).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
