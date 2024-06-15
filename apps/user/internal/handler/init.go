package controller

import (
	"github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
	pb.UserServiceServer
}

func NewController(db *gorm.DB) *Handler {
	return &Handler{db: db}
}
