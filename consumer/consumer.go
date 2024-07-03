package main

import (
	"encoding/json"
	"fmt"

	"github.com/RYANCOAL9999/DisSystem/tools"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Params struct {
	Username string
}

func Consume(queueName string, channel *amqp.Channel) {
	// Consume the response from the response queue
	messages, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)

			var params Params
			err := json.Unmarshal(message.Body, &params)
			if err != nil {
				log.Printf("Error decoding message body: %v\n", err)
				continue
			}

			// Initialize the database
			var database *tools.DatabaseInterface
			database, err = tools.NewDatabase()
			if err != nil {
				log.Printf("Error initializing database: %v\n", err)
				continue
			}

			// Get user coin details
			var tokenDetails = (*database).GetUserCoins(params.Username)
			if tokenDetails == nil {
				log.Printf("Error retrieving user coins for username: %s\n", params.Username)
				continue
			}

			// Log the token details
			log.Printf("User: %s, Coins: %d\n", params.Username, tokenDetails.Coins)
		}
	}()

	<-forever
}

func RabbitMQConnect() (*amqp.Connection, *amqp.Channel, error) {

	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	// Establishing a connection to RabbitMQ
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	fmt.Println("Successfully connected to RabbitMQ instance")

	// Opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		connection.Close() // Ensure connection is closed if channel creation fails
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	fmt.Println("Successfully opened a channel")

	return connection, channel, nil
}

func main() {

	connection, channel, err := RabbitMQConnect()

	if err != nil {
		log.Fatalf("RabbitMQ connection error: %v", err)
	}

	defer connection.Close()

	defer channel.Close()

	// Start consuming messages
	go Consume("coin_balance_requests", channel)

	// Keep the main function running indefinitely
	select {}

}
