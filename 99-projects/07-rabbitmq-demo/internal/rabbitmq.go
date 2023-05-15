package internal

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	// TCP connection used by client
	Conn *amqp.Connection

	// AMQP channel. TCP connection can have multiple channels
	Ch *amqp.Channel
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost)

	return amqp.Dial(url)
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()

	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{Conn: conn, Ch: ch}, nil
}

func (rc RabbitClient) Close() error {
	return rc.Ch.Close()
}

func (rc RabbitClient) CreateQueue(queueName string, durable, autodelete bool) error {
	_, err := rc.Ch.QueueDeclare(queueName, durable, autodelete, false, false, nil)
	return err
}

func (rc RabbitClient) CreateBinding(name, binding, exchange string) error {
	return rc.Ch.QueueBind(name, binding, exchange, false, nil)
}

func (rc RabbitClient) Send(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	return rc.Ch.PublishWithContext(ctx, exchange, routingKey, true, false, options)
}

func (rc RabbitClient) ReceiveChan(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return rc.Ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}
