package controller

import (
	"context"

	"github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := c.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
