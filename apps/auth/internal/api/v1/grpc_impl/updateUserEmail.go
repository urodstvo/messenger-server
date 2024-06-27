package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateUserEmail(ctx context.Context, req *proto.UpdateUserEmailRequest) (*emptypb.Empty, error) {
	user := &models.User{ID: req.UserId}

	if err := i.db.WithContext(ctx).Model(user).Update("email", req.Email).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
