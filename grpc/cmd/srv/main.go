package main

import (
	"context"
	"log"
	"net"

	"grpc/proto/pb"

	"google.golang.org/grpc"
)

// Implémentation du service gRPC
type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserByID(ctx context.Context, req *pb.IDRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name: "John Doe",
		Age:  30,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Échec de l'écoute: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Échec du serveur gRPC: %v", err)
	}
}
