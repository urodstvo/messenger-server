package implementation

import (
	"context"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateMessageText(ctx context.Context, req *proto.UpdateMessageTextRequest) (*emptypb.Empty, error) {
	message := &models.Message{ID: req.MessageId}
	if err := i.db.WithContext(ctx).Model(message).Update("text", req.Text).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
