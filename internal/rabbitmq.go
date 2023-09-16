package internal

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn *amqp.Connection
	// Channel is used to process / Seng messages
	ch *amqp.Channel
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rC *RabbitClient) Close() error {
	return rC.ch.Close()
}

func (rC *RabbitClient) CreateQueue(queueName string, durable, autoDelete bool) error {
	_, err := rC.ch.QueueDeclare(queueName, durable, autoDelete, false, false, nil)
	return err
}
