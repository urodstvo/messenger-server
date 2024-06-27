package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateUserFirstLastName(ctx context.Context, req *proto.UpdateUserFirstLastNameRequest) (*emptypb.Empty, error) {
	user := &models.User{
		ID: req.UserId,
	}

	if err := i.db.WithContext(ctx).Model(user).Updates(map[string]interface{}{"first_name": req.FirstName, "last_name": req.LastName}).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
