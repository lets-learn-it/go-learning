package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func callSayHelloServerStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Streaming started...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("failed. Err: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failure while receiving greet. Err: %v", err)
		}

		log.Println(message.Message)
	}
	log.Println("Streaming finished.")
}
