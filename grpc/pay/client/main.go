package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "srpc/pay/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ProcessPayment(ctx, &pb.PaymentRequest{
		Amount:     100,
		OutTradeNo: "123456",
	})
	if err != nil {
		log.Printf("grpc status: %d", status.Convert(err).Code())
		log.Printf("grpc status: %v", status.Convert(err).Message())
		log.Fatalf("could not Pay: %v", err)
	}
	log.Printf("outTradeNo: %s, status: %s, transactionId: %s", r.OutTradeNo, r.Status, r.TransactionId)

	// 打印grpc状态码
	log.Printf("grpc status: %v", r.GetStatus())
}
