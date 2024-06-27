package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetUsersLikeFirstLastName(ctx context.Context, req *proto.GetUsersLikeFirstLastNameRequest) (*proto.GetUsersLikeFirstLastNameResponse, error) {
	var users []models.User

	if err := i.db.WithContext(ctx).Where("first_name LIKE ? AND last_name LIKE ?", "%"+req.FirstName+"%", "%"+req.LastName+"%").Find(&users).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	var usersResponse []*proto.User
	for _, user := range users {
		usersResponse = append(usersResponse, &proto.User{
			Id:        user.ID,
			Email:     user.Email,
			Tag:       user.Tag,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}

	return &proto.GetUsersLikeFirstLastNameResponse{Users: usersResponse}, nil

}
