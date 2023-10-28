package main

import (
	"context"
	"log"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	log.Println("Hello")
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
