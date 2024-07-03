package api

import "github.com/streadway/amqp"

// Coint Balance Params
type Params struct {
	Username string
}

type CoinBalanceResponse struct {
	//Success Code, Ususally 200
	Code int
}

type QyqueueType struct {
	Durable   bool
	Delete    bool
	Exclusive bool
	No_wait   bool
	Args      amqp.Table
}
