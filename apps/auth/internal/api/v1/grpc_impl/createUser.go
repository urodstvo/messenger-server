package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*emptypb.Empty, error) {
	user := &models.User{
		Email: req.Email,
		Tag:   req.Tag,
	}

	if err := i.db.WithContext(ctx).Create(user).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
