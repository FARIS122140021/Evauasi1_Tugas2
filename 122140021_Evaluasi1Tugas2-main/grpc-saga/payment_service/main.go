package main

import (
    "context"
    "log"
    "net"

    pb "grpc-saga/proto"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
    log.Printf("Payment processed for order: %s", req.OrderId)
    return &pb.ProcessPaymentResponse{OrderId: req.OrderId, Status: "SUCCESS"}, nil
}

func (s *server) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
    log.Printf("Payment refunded for order: %s", req.OrderId)
    return &pb.RefundPaymentResponse{Message: "Payment refunded"}, nil
}

func main() {
    listener, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterPaymentServiceServer(s, &server{})

    log.Println("Payment Service running on port 50052")
    if err := s.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
