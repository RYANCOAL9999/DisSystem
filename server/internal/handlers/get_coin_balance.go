package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/server/api"
	"github.com/RYANCOAL9999/DisSystem/server/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

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

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)

			var params tools.CoinDetails
			err := json.Unmarshal(message.Body, &params)

			if err != nil {
				log.Printf("Error decoding message body: %v\n", err)
				continue
			}

			// Log the token details
			log.Printf("User: %s, Coins: %d\n", params.Username, params.Coins)

			// Store the message
			api.MessagesLock.Lock()
			api.Messages = append(api.Messages, params)
			api.MessagesLock.Unlock()
		}
	}()

	<-forever
}

func GetCoinBalance(queueName string, writer http.ResponseWriter, request *http.Request, channel *amqp.Channel) {

	params := api.Params{}
	var err error

	err = schema.NewDecoder().Decode(&params, request.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	// Call the external service
	externalURL := fmt.Sprintf("http://localhost:3000/%s/coins", params.Username)
	resp, err := http.Get(externalURL)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	response := api.OkayResponse{Code: resp.StatusCode}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}
}
