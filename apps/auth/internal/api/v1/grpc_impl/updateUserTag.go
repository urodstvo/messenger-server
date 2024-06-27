package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateUserTag(ctx context.Context, req *proto.UpdateUserTagRequest) (*emptypb.Empty, error) {
	if err := i.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", req.UserId).Update("tag", req.Tag).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
