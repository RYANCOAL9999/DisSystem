package handlers

import (
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/server/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/streadway/amqp"
)

func Handler(r *chi.Mux, channel *amqp.Channel) {

	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/hearts", GetUserHearts)

		router.Get("/messages", PollingHandler)

		router.Get("/coinShow", func(writer http.ResponseWriter, request *http.Request) {
			GetCoinBalance("myQueue", writer, request, channel)
		})
	})
}
