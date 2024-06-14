package main

import (
	"fmt"
	"log"
	"net"

	database "github.com/urodstvo/messenger-server/apps/chat/db"
	controller "github.com/urodstvo/messenger-server/apps/chat/internal/handler"
	"google.golang.org/grpc"

	pb "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__/pb"
)

func main() {
	// init database connection
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalln("@Chat Service: Error while connecting to database: ", err.Error())
	}

	handlers := controller.NewController(db)

	// Create a TCP Listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8000"))
	if err != nil {
		log.Fatalln("Could not start TCP server: ", err.Error())
	}

	// Create a gRPC Server and Register the Definitions
	grpcSrv := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcSrv, handlers)

	// Log a quick message
	log.Println("Starting Chat Service...")

	// Run the Server
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalln("Critical Error in Chat Service: ", err.Error())
	}
}
