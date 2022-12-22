package main

import (
	"context"
	"log"

	pb "github.com/radulucut/go-grpc-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:5001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	client := pb.NewChatClient(conn)

	message, err := client.Send(context.Background(), &pb.Message{Text: "Hello world!"})

	if err != nil {
		log.Fatalf("Error when calling Send: %v", err)
	}

	log.Printf("Response from server: %v", message.GetText())
}
