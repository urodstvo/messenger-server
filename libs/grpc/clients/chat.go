package clients

import (
	"log"

	"github.com/urodstvo/messenger-server/libs/grpc/constants"
	pb "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"google.golang.org/grpc"
)

func NewChatAppClient(isProduction bool) pb.ChatServiceClient {
	serverAddress := createClientAddr(isProduction, "chat", constants.CHAT_SERVER_PORT)

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
