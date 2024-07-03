package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/publisher/api"
	"github.com/go-chi/chi"
	"github.com/streadway/amqp"
)

func QueueDeclare(channel *amqp.Channel, name string, qt api.QyqueueType) (string, error) {
	queue, err := channel.QueueDeclare(
		name,
		qt.Durable,
		qt.Delete,
		qt.Exclusive,
		qt.No_wait,
		qt.Args,
	)
	if err != nil {
		return "", fmt.Errorf("failed to declare queue: %w", err)
	}

	return queue.Name, nil
}

func Handler(r *chi.Mux, channel *amqp.Channel) {

	// Define the queue parameters
	queueName, err := QueueDeclare(channel, "myQueue", api.QyqueueType{})

	if err != nil {
		log.Fatalf("Queue declare error: %v", err)
	}

	r.Route("/{username}", func(router chi.Router) {
		router.Get("/coins", func(writer http.ResponseWriter, request *http.Request) {
			GetCoinBalance(queueName, writer, request, channel)
		})
	})

}
