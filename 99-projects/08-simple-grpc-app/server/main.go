package main

import (
	"log"
	"net"

	"github.com/lets-learn-it/simple-grpc-app/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on %s. Err: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve. Err: %v", err)
	}

}
