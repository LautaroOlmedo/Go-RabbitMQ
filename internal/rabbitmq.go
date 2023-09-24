package internal

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitClient struct {
	// The connection used by the client
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
	if err := ch.Confirm(false); err != nil {
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

// CreateQueue will create a new queue based on given configuration
func (rC *RabbitClient) CreateQueue(queueName string, durable, autoDelete bool) error {
	_, err := rC.ch.QueueDeclare(queueName, durable, autoDelete, false, false, nil)
	return err
}

// CreateBinding will bind the current channel to the given exchange using the routingKey provided
func (rC *RabbitClient) CreateBinding(queueName, binding, exchange string) error {
	return rC.ch.QueueBind(queueName, binding, exchange, false, nil)
}

func (rC *RabbitClient) Send(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	// Mandatory is used to determine if an error should be returned upon failure
	confirmation, err := rC.ch.PublishWithDeferredConfirmWithContext(ctx, exchange, routingKey, true, false, options)

	if err != nil {
		return err
	}
	log.Println(confirmation.Wait())
	return nil
}

// Consume is used to consume a queue
func (rC *RabbitClient) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return rC.ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}

func (rC *RabbitClient) ApplyQos(count, size int, global bool) error {
	return rC.ch.Qos(count, size, global)
}
