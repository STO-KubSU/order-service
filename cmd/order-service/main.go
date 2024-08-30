package main

import (
	"log"
	"net"

	"github.com/STO-KubSU/order-service/internal/service"
	pb "github.com/STO-KubSU/orderpb"
	userpb "github.com/STO-KubSU/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Подключаемся к User Service
	conn, err := grpc.Dial("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := userpb.NewUserServiceClient(conn)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderService(userClient))

	log.Println("Order service is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
