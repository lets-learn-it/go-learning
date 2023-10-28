package main

import (
	"fmt"
	"log"

	"github.com/lets-learn-it/simple-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = 8080
	host = "localhost"
)

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect. Err: %v", err)
	}

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)
	names := &proto.NamesList{
		Names: []string{"Ram", "Sita", "Krishna", "Rukmini"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBiStream(client, names)
}
