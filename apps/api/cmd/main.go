package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/urodstvo/messenger-server/libs/grpc/clients"
	pb "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
)

type HandlerController struct {
	chats pb.ChatServiceClient
	users pb.UserServiceClient
}

func NewHandlerController() (*HandlerController, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create the Client
	chatClient := clients.NewChatAppClient(false, 8000)
	userClient := clients.NewUserAppClient(false, 8001)

	// Return the Controller
	return &HandlerController{
		chats: chatClient,
		users: userClient,
	}, nil
}

func main() {
	handler, err := NewHandlerController()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
		_, err := handler.users.CreateUser(context.Background(), &pb.CreateUserRequest{Username: "John", Password: "123"})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		_, err = handler.users.CreateUser(context.Background(), &pb.CreateUserRequest{Username: "Jane", Password: "123"})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		_, err = handler.chats.CreateChat(context.Background(), &pb.CreateChatRequest{Name: "John & Jane", IsPublic: true, Participants: []uint32{}})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})

	http.ListenAndServe(":80", nil)
}
