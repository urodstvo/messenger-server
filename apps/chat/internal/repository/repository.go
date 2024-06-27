package repository

import (
	"context"

	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *models.Chat) (uint32, error)
	GetById(ctx context.Context, id uint32) (*models.Chat, error)
	GetByTag(ctx context.Context, tag string) (*models.Chat, error)
	Update(ctx context.Context, chat *models.Chat, columns ...string) error
	Delete(ctx context.Context, id uint32) error
}

type MessageRepository interface {
	Create(ctx context.Context, message *models.Message) (uint32, error)
	GetById(ctx context.Context, id uint32) (*models.Message, error)
	GetByChatId(ctx context.Context, id uint32) (*[]models.Message, error)
	Update(ctx context.Context, message *models.Message, columns ...string) error
	Delete(ctx context.Context, id uint32) error
}

type ParticipantRepository interface {
	Create(ctx context.Context, participant *models.Participant) error
	GetByUserId(ctx context.Context, id uint32) (*[]models.Participant, error)
	GetByChatId(ctx context.Context, id uint32) (*[]models.Participant, error)
	Update(ctx context.Context, participant *models.Participant) error
	Delete(ctx context.Context, chatId, userId uint32) error
}
