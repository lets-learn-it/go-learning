package main

import (
	"context"
	"log"
	"time"

	"avabodha.in/rabbitmq-demo/internal"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("pskp", "secret", "localhost:5672", "customers")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.CreateQueue("customer_created", true, false); err != nil {
		panic(err)
	}

	if err := client.CreateQueue("customer_test", false, true); err != nil {
		panic(err)
	}

	if err := client.CreateBinding("customer_created", "customers.created.*", "customer_events"); err != nil {
		panic(err)
	}

	// start consuming messages
	channel, err := client.ReceiveChan("customer_created", "worker", false)

	if err != nil {
		panic(err)
	}

	go func() {
		for message := range channel {
			log.Println("Message: ", string(message.Body))

			if err := message.Ack(false); err != nil {
				log.Println("Acknowledge message failed")
				continue
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Send(ctx, "customer_events", "customers.created.us", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Persistent,
		Body:         []byte(`An cool message between services!!!`),
	}); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
}
