package main

import (
	"log"
	"rabbitMQ/internal"
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

	messageBus, err := client.Consume("customers_created", "email-service", false)

	if err != nil {
		panic(err)
	}

	var blocking chan struct{}

	go func() {
		for msg := range messageBus {
			log.Printf("New Message %v \n", msg)
			if err := msg.Ack(false); err != nil {
				log.Printf("Acknowledge message failed")
				continue
			}
			log.Printf("Acknowledge message %s\n", msg.MessageId)
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")

	<-blocking
}
