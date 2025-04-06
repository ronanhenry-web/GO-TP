package main

import (
	"context"
	"grpc/proto/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	gRPCClient := pb.NewUserServiceClient(conn)

	userResponse, err := gRPCClient.GetUserByID(context.Background(), &pb.IDRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	log.Printf("User: %s, Age: %d", userResponse.Name, userResponse.Age)
}
