package implementation

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"github.com/urodstvo/messenger-server/libs/jwt"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm"
)

func (i *Implementation) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	otp := &models.OTP{
		Email: req.Email,
	}

	if err := i.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).First(otp).Error; err != nil {
			i.logger.Error(err.Error())
			return err
		}

		if otp.Otp != req.Otp {
			return errors.New("wrong otp")
		}

		if otp.ExpiresAt.Before(time.Now()) {
			return errors.New("otp expired")
		}

		if err := tx.WithContext(ctx).Delete(otp).Error; err != nil {
			i.logger.Error(err.Error())
			return err
		}

		user := &models.User{
			Email: req.Email,
			Tag:   req.Tag,
		}

		if req.FirstName != nil {
			user.FirstName = *req.FirstName
		}

		if req.LastName != nil {
			user.LastName = *req.LastName
		}

		if err := tx.WithContext(ctx).Create(user).Error; err != nil {
			i.logger.Error(err.Error())
			return err
		}

		return nil
	}); err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	var user models.User
	if err := i.db.WithContext(ctx).First(&user, "email = ?", req.Email).Error; err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	access_token, err := jwt.Generate(fmt.Sprint(user.ID), time.Now().Add(20*time.Minute), i.cfg.JWTSecret)
	if err != nil {
		i.logger.Error("generate access token error:", err)
		return nil, err
	}

	refresh_token, err := jwt.Generate(fmt.Sprint(user.ID), time.Now().Add(7*24*time.Hour), i.cfg.JWTSecret)
	if err != nil {
		i.logger.Error("generate refresh token error:", err)
		return nil, err
	}

	return &proto.RegisterResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}
