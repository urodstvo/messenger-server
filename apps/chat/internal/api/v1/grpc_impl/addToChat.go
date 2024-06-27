package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddParticipantToChat(ctx context.Context, req *proto.AddParticipantRequest) (*emptypb.Empty, error) {
	participant := &models.Participant{
		ChatID: req.ChatId,
		UserID: req.UserId,
	}

	if err := i.db.WithContext(ctx).Create(participant).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
