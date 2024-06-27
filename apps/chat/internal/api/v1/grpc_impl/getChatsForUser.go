package implementation

import (
	"context"

	"github.com/urodstvo/messenger-server/apps/chat/internal/converter"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (i *Implementation) GetChatsForUser(ctx context.Context, req *proto.GetChatsForUserRequest) (*proto.GetChatsForUserResponse, error) {
	var chatIds []uint32
	var parts []models.Participant
	if err := i.db.WithContext(ctx).Where("user_id = ?", req.UserId).Find(&parts).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	for _, part := range parts {
		chatIds = append(chatIds, part.ChatID)
	}

	var chats []models.Chat
	if err := i.db.WithContext(ctx).Preload(clause.Associations).Where("id in ?", chatIds).Find(&chats).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	var chatsResponse []*proto.Chat

	for _, chat := range chats {
		chatsResponse = append(chatsResponse, converter.FromChatModelToProto(&chat))
	}

	return &proto.GetChatsForUserResponse{Chats: chatsResponse}, nil

}
