package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/api"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetCoinBalance(queueName string, writer http.ResponseWriter, request *http.Request, channel *amqp.Channel) {

	var params api.Params
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, request.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
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
		api.InternalErrorHandler(writer)
		return
	}

}
