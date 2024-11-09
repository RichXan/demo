package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "xrpc/pay/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement pay.PaymentServiceServer
type server struct {
	pb.UnimplementedPaymentServiceServer
}

// ProcessPayment implements pay.PaymentServiceServer
func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Here you would add logic to process the payment
	// For example, call a third-party payment gateway API

	// Simulate a successful payment process
	if req.Amount <= 0 {
		return nil, fmt.Errorf("invalid amount")
	}

	time.Sleep(time.Second * 3)

	// 判断ctx.Deadline
	if deadline, ok := ctx.Deadline(); ok {
		log.Printf("deadline: %v", deadline)
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	}

	return &pb.PaymentResponse{
		TransactionId: "1234567890",
		Status:        "Success",
	}, nil
}

// Implement other RPC methods similarly...

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
