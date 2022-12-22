package main

import (
	"context"
	"log"
	"net"

	pb "github.com/radulucut/go-grpc-chat/proto"
	"google.golang.org/grpc"
)

const address = ":5001"

type server struct {
	pb.UnimplementedChatServer
}

func (s *server) Send(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Received message from client: %v", in.GetText())

	return &pb.Message{Text: "Message received!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on %v: %v", address, err)
	}

	srv := grpc.NewServer()

	pb.RegisterChatServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v: %v", address, err)
	}
}
