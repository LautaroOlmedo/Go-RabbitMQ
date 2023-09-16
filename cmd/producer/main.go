package main

import (
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

	time.Sleep(20 * time.Second)
	log.Println(client)
}
