package clients

import (
	"log"

	"github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	"google.golang.org/grpc"
)

func NewChatAppClient(isProduction bool, port uint) pb.ChatServiceClient {
	serverAddress := createClientAddr(isProduction, "chat", int(port))

	conn, err := grpc.NewClient(
		serverAddress,
		defaultClientsOptions...,
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewChatServiceClient(conn)

	return c
}
