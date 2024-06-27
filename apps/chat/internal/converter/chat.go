package converter

import (
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/timestamppb"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
)

func FromChatModelToProto(chat *models.Chat) *proto.Chat {
	chatParticipants := make([]*proto.Participant, len(chat.Participants))
	chatMessages := make([]*proto.Message, len(chat.Messages))

	for i, participant := range chat.Participants {
		chatParticipants[i] = &proto.Participant{
			UserId: participant.UserID,
			ChatId: participant.ChatID,
			Role:   proto.ParticipantRole(proto.ParticipantRole_value[string(participant.Role)]),
		}
	}

	for i, message := range chat.Messages {
		chatMessages[i] = &proto.Message{
			Id:        message.ID,
			ChatId:    message.ChatID,
			AuthorId:  message.AuthorID,
			Text:      message.Text,
			Status:    proto.MessageStatus(proto.MessageStatus_value[string(message.Status)]),
			CreatedAt: timestamppb.New(message.CreatedAt),
			UpdatedAt: timestamppb.New(message.UpdatedAt),
		}
	}

	chatResponse := &proto.Chat{
		Id:           chat.ID,
		Name:         chat.Name,
		Tag:          chat.Tag,
		IsPublic:     chat.IsPublic,
		Participants: chatParticipants,
		Messages:     chatMessages,
		CreatedAt:    timestamppb.New(chat.CreatedAt),
		UpdatedAt:    timestamppb.New(chat.UpdatedAt),
	}

	return chatResponse
}
