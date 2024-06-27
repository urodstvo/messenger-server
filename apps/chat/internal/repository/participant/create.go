package participant

import (
	"context"

	errors "github.com/urodstvo/messenger-server/apps/chat/internal/repository"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (r *repository) Create(ctx context.Context, participant *models.Participant) error {
	if err := r.db.WithContext(ctx).Create(&participant).Error; err != nil {
		return errors.NewRepositoryError(err)
	}
	return nil
}
