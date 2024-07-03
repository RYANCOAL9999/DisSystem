package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/streadway/amqp"
)

type QyqueueType struct {
	durable   bool
	delete    bool
	exclusive bool
	no_wait   bool
	args      amqp.Table
}

func QueueDeclare(channel *amqp.Channel, name string, qt QyqueueType) (string, error) {
	queue, err := channel.QueueDeclare(
		name,         // name
		qt.durable,   // durable
		qt.delete,    // auto delete
		qt.exclusive, // exclusive
		qt.no_wait,   // no-wait
		qt.args,      // args
	)
	if err != nil {
		return "", fmt.Errorf("failed to declare queue: %w", err)
	}

	return queue.Name, nil
}

func Handler(r *chi.Mux, channel *amqp.Channel) {

	// Define the queue parameters
	queueName, err := QueueDeclare(channel, "myQueue", QyqueueType{})

	if err != nil {
		log.Fatalf("Queue declare error: %v", err)
	}

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Get("/coin", func(writer http.ResponseWriter, request *http.Request) {
			GetCoinBalance(queueName, writer, request, channel)
		})
	})
}
