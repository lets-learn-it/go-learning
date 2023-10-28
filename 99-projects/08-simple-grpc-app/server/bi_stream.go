package main

import (
	"io"
	"log"

	"github.com/lets-learn-it/simple-grpc-app/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream proto.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(req.Name)
		err = stream.Send(&proto.HelloResponse{
			Message: "Hello " + req.Name,
		})

		if err != nil {
			log.Println(err)
			return err
		}
	}
}
