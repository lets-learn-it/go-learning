package main

import (
	"io"
	"log"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.MessageList{Message: messages})
		}
		if err != nil {
			log.Fatalf("failure. Err: %v", err)
		}
		log.Println(req.Name)
		messages = append(messages, req.Name)
	}
}
