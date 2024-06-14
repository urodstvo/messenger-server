package controller

import (
	"context"

	"github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Handler) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*emptypb.Empty, error) {
	participants := make([]models.Participant, len(req.Participants))
	for user := range participants {
		participants[user] = models.Participant{UserID: req.Participants[user]}
	}

	chat := &models.Chat{
		Name:         req.Name,
		IsPublic:     req.IsPublic,
		Participants: participants,
	}

	if err := c.db.WithContext(ctx).Create(chat).Error; err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
