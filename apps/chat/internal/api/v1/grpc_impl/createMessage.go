package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (i *Implementation) CreateMessage(ctx context.Context, req *proto.CreateMessageRequest) (*proto.CreateMessageResponse, error) {
	message := &models.Message{
		ChatID:   req.ChatId,
		AuthorID: req.AuthorId,
		Text:     req.Text,
	}

	if err := i.db.WithContext(ctx).Create(message).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &proto.CreateMessageResponse{Id: message.ID}, nil

}
