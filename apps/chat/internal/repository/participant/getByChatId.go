package participant

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (r *repository) GetByChatId(ctx context.Context, id uint32) (*[]models.Participant, error) {
	var participants []models.Participant
	if err := r.db.WithContext(ctx).Preload(clause.Associations).Where("chat_id = ?", id).Find(&participants).Error; err != nil {
		return nil, errors.NewRepositoryError(err)
	}

	return &participants, nil
}
