package main

import (
	"context"
	"log"
	"time"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func callSayHelloClientStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	for _, name := range names.Names {
		err = stream.Send(&proto.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Err: %v", err)
		}

		log.Println("sent ", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	log.Printf("%v", res)
}
