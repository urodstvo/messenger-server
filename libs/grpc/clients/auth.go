package clients

import (
	"log"

	"github.com/urodstvo/messenger-server/libs/grpc/constants"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"google.golang.org/grpc"
)

func NewAuthAppClient(isProduction bool) proto.AuthServiceClient {
	serverAddress := createClientAddr(isProduction, "user", constants.AUTH_SERVER_PORT)

	conn, err := grpc.NewClient(
		serverAddress,
		defaultClientsOptions...,
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewAuthServiceClient(conn)

	return c
}
