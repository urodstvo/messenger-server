package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (i *Implementation) CheckEmailExists(ctx context.Context, req *proto.CheckEmailExistsRequest) (*proto.CheckEmailExistsResponse, error) {
	var user models.User

	result := i.db.WithContext(ctx).Where("email = ?", req.Email).Find(&user)

	return &proto.CheckEmailExistsResponse{Exists: result.RowsAffected > 0}, nil

}
