package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteChatById(ctx context.Context, req *proto.DeleteChatByIdRequest) (*emptypb.Empty, error) {
	if err := i.db.WithContext(ctx).Delete(&models.Chat{}, req.ChatId).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
