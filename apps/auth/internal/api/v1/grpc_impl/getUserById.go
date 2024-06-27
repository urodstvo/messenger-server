package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetUserById(ctx context.Context, req *proto.GetUserByIdRequest) (*proto.GetUserByIdResponse, error) {
	user := &models.User{ID: req.UserId}
	if err := i.db.WithContext(ctx).First(&user).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	userResponse := &proto.User{
		Id:        user.ID,
		Email:     user.Email,
		Tag:       user.Tag,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return &proto.GetUserByIdResponse{User: userResponse}, nil

}
