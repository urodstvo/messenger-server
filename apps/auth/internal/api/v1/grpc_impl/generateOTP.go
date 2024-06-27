package implementation

import (
	"context"
	"log/slog"
	"time"

	"github.com/urodstvo/messenger-server/apps/auth/internal/otp"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GenerateOTP(ctx context.Context, req *proto.GenerateOTPRequest) (*emptypb.Empty, error) {
	model := &models.OTP{
		Email:     req.Email,
		Otp:       otp.Generate(),
		ExpiresAt: time.Now().Add(2 * time.Minute),
	}

	if err := i.db.WithContext(ctx).Create(model).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	if err := i.email.SendOTP([]string{model.Email}, model.Otp); err != nil {
		i.logger.Error("failed to send otp", slog.Any("err", err))
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
