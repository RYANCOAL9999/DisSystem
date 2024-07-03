package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/streadway/amqp"
)

type qyqueueType struct {
	durable   bool
	delete    bool
	exclusive bool
	no_wait   bool
	args      amqp.Table
}

func queueDeclare(channel *amqp.Channel, name string, qt qyqueueType) (string, error) {
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
	queueName, err := queueDeclare(channel, "myQueue", qyqueueType{})

	if err != nil {
		log.Fatalf("Queue declare error: %v", err)
	}

	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/hearts", GetUserHearts)

		router.Get("/coinShow", func(writer http.ResponseWriter, request *http.Request) {
			GetCoinBalance(queueName, writer, request, channel)
		})
	})
}
