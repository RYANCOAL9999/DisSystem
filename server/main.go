package main

import (
	"fmt"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

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

	log.SetReportCaller(true)

	connection, channel, err := RabbitMQConnect()

	if err != nil {
		log.Fatalf("RabbitMQ connection error: %v", err)
	}

	defer connection.Close()

	defer channel.Close()

	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r, channel)

	fmt.Println("Starting Go API service...")

	fmt.Println(`
	______     ______        ______     ______   __    
   /\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
   \ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
	\ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
	 \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	listenErr := http.ListenAndServe("localhost:8080", r)

	if listenErr != nil {
		log.Error(listenErr)
	}

}
