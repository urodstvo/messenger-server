package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteMessageById(ctx context.Context, req *proto.DeleteMessageByIdRequest) (*emptypb.Empty, error) {
	if err := i.db.WithContext(ctx).Delete(&models.Message{}, req.MessageId).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
