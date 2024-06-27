package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetMessagesForChat(ctx context.Context, req *proto.GetMessagesForChatRequest) (*proto.GetMessagesForChatResponse, error) {
	chat := &models.Chat{ID: req.ChatId}
	if err := i.db.WithContext(ctx).Preload("Messages").Find(chat).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	messages := chat.Messages
	var messagesResponse []*proto.Message

	for _, message := range messages {
		messagesResponse = append(messagesResponse, &proto.Message{
			Id:        message.ID,
			ChatId:    message.ChatID,
			AuthorId:  message.AuthorID,
			Text:      message.Text,
			Status:    proto.MessageStatus(proto.MessageStatus_value[string(message.Status)]),
			CreatedAt: timestamppb.New(message.CreatedAt),
			UpdatedAt: timestamppb.New(message.UpdatedAt),
		})
	}

	return &proto.GetMessagesForChatResponse{Messages: messagesResponse}, nil

}
