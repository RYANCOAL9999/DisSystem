package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/publisher/api"
	"github.com/RYANCOAL9999/DisSystem/publisher/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Publish(queueName string, channel *amqp.Channel, params *tools.CoinDetails) {

	// Initialize the database
	requestBody, err := json.Marshal(params)
	if err != nil {
		log.Error(err)
		return
	}

	// Publish the request to the queue
	err = channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        requestBody,
		})

	if err != nil {
		log.Error(err)
		return
	}

}

func GetCoinBalance(queueName string, writer http.ResponseWriter, request *http.Request, channel *amqp.Channel) {

	var username string = chi.URLParam(request, "username")
	response := api.CoinBalanceResponse{Code: http.StatusOK}
	var err error

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(response)

	if err != nil {
		log.Error(err)
		return
	}

	// Initialize the database
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return
	}

	var tokenDetails *tools.CoinDetails = (*database).GetUserCoins(username)
	if tokenDetails == nil {
		log.Printf("Error retrieving user coins for username: %s\n", username)
		return
	}

	Publish(queueName, channel, tokenDetails)

}
