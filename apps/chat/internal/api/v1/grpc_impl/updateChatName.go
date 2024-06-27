package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateChatName(ctx context.Context, req *proto.UpdateChatRequest) (*emptypb.Empty, error) {
	chat := &models.Chat{ID: req.ChatId}
	if err := i.db.WithContext(ctx).Model(chat).Update("name", req.Name).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
