// order-service/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "path/to/order-service"
)

type orderServer struct{}

func (s *orderServer) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	orderID := generateID()
	order := &pb.OrderResponse{
		Id:     orderID,
		Amount: req.Amount,
		Status: "Pending",
	}
	// Save order details or communicate with a database

	return order, nil
}

func (s *orderServer) GetOrder(ctx context.Context, query *pb.OrderQuery) (*pb.OrderResponse, error) {
	// Retrieve order details by ID or communicate with a database
	order := &pb.OrderResponse{
		Id:     query.OrderId,
		Amount: 42.0, // Replace with actual order amount
		Status: "Delivered", // Replace with actual order status
	}

	return order, nil
}

func generateID() string {
	// Implement your own ID generation logic
	return "12345"
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &orderServer{})

	fmt.Println("Order Service listening on :8080...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
