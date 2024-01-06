// payment-service/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "path/to/payment-service"
)

type paymentServer struct{}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Simulate payment processing logic
	payment := &pb.PaymentResponse{
		OrderId: req.OrderId,
		Amount:  req.Amount,
		Status:  "Success",
	}
	// Communicate with the Order Service to update order status
	updateOrderStatus(req.OrderId, "Paid")

	return payment, nil
}

func updateOrderStatus(orderID, status string) {
	// Implement communication logic with the Order Service to update order status
	fmt.Printf("Updating order %s status to %s\n", orderID, status)
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterPaymentServiceServer(server, &paymentServer{})

	fmt.Println("Payment Service listening on :8081...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
