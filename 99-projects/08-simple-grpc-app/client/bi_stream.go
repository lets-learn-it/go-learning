package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func callSayHelloBiStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("bidirectional streaming started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Err: %v", err)
			}

			log.Println(message.Message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Err: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("streaming ended")
}
