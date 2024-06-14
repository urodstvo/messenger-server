package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HandlerController struct {
	chats pb.ChatServiceClient
}

func NewHandlerController() (*HandlerController, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Get the Accounts Service Host
	chatURI, exists := os.LookupEnv("CHAT_SERVICE_URI")
	if !exists {
		return nil, fmt.Errorf("no CHAT_SERVICE_URI specified")
	}

	// Connect to the Service
	chatConn, err := grpc.NewClient(chatURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create the Client
	chatClient := pb.NewChatServiceClient(chatConn)

	// Return the Controller
	return &HandlerController{
		chats: chatClient,
	}, nil
}

func main() {
	handler, err := NewHandlerController()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
		_, err := handler.chats.CreateChat(context.Background(), &pb.CreateChatRequest{Name: "John", IsPublic: true, Participants: []uint32{1, 2}})

		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
