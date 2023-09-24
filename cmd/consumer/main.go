package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"

	"rabbitMQ/internal"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("lautaro", "secret", "localhost:5672", "customers")
	if err != nil {
		log.Fatalf("error when connecting to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)

	if err != nil {
		log.Fatalf("error when creating the RabbitMQ client: %v", err)
		return
	}
	defer client.Close()

	messageBus, err := client.Consume("customers_created", "email-service", false)

	if err != nil {
		log.Fatalf("error while consuming messages: %v", err)
		return
	}

	var blocking chan struct{}

	// set a timeout for 15 secs
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// apply a hard limit on the Server
	if err := client.ApplyQos(10, 0, true); err != nil {
		log.Fatalf("error applying strict limit to server: %v", err)
		return
	}

	// errgroup allows us concurrent tasks
	g.SetLimit(10)
	go func() {
		for message := range messageBus {
			// spawn a worker
			msg := message
			g.Go(func() error {
				log.Printf("New message: %v", msg)
				time.Sleep(10 * time.Second)
				if err := msg.Ack(false); err != nil {
					log.Println("Ack message failed")
					return err
				}
				log.Printf("Acknowledged message %s\n", message.MessageId)
				return nil
			})
		}
	}()
	log.Println("Consuming, to close the program press CTRL+C")

	<-blocking
}
