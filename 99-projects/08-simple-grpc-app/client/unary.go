package main

import (
	"context"
	"log"
	"time"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func callSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Client failed. Err: %v", err)
	}

	log.Println(res.Message)
}
