package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitMQ/internal"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("lautaro", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)

	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.CreateQueue("customers_created", true, false); err != nil {
		panic(err)
	}

	if err := client.CreateQueue("customers_test", false, true); err != nil {
		panic(err)
	}

	if err := client.CreateBinding("customers_created", "customers.created.*", "customer_events"); err != nil {
		panic(err)
	}

	if err := client.CreateBinding("customers_test", "customers.*", "customer_events"); err != nil {
		panic(err)
	}

	myContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Send(myContext, "customer_events", "customers.created.us", amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         []byte(`Lautaro`),
	}); err != nil {
		panic(err)
	}

	time.Sleep(20 * time.Second)
	log.Println(client)
}
