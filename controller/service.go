package controller

import (
	"context"
	pb "grpc_ex/protobuf"
	"log"
)

// service -
type service struct{}

// NewServer -
func NewServer() *service {
	return &service{}
}

// SayHello -
func (s *service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("SayHello Service Received: %v\n", in.GetName())

	return &pb.HelloResponse{
		Reply: "Hello, " + in.GetName(),
	}, nil
}
