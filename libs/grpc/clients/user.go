package clients

import (
	"log"

	"github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	"google.golang.org/grpc"
)

func NewUserAppClient(isProduction bool, port uint) pb.UserServiceClient {
	serverAddress := createClientAddr(isProduction, "user", int(port))

	conn, err := grpc.NewClient(
		serverAddress,
		defaultClientsOptions...,
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c
}
