package implementation

import (
	"context"

	"github.com/urodstvo/messenger-server/apps/chat/internal/converter"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (i *Implementation) GetChatById(ctx context.Context, req *proto.GetChatByIdRequest) (*proto.GetChatByIdResponse, error) {
	chat := &models.Chat{ID: req.ChatId}

	if err := i.db.WithContext(ctx).Preload(clause.Associations).First(chat).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	chatResponse := converter.FromChatModelToProto(chat)

	return &proto.GetChatByIdResponse{Chat: chatResponse}, nil

}
