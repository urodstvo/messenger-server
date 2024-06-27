package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (i *Implementation) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {
	participants := make([]models.Participant, len(req.Participants))

	for i, userId := range req.Participants {
		participants[i] = models.Participant{
			UserID: userId,
		}
	}

	chat := &models.Chat{
		Name:         req.Name,
		Tag:          req.Tag,
		IsPublic:     req.IsPublic,
		Participants: participants,
	}

	if err := i.db.WithContext(ctx).Create(chat).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &proto.CreateChatResponse{Id: chat.ID}, nil

}
