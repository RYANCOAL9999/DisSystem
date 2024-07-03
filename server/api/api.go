package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/RYANCOAL9999/DisSystem/server/internal/tools"
)

// Coint Balance Params
type Params struct {
	Username string
}

type OkayResponse struct {

	//Success Code, Ususally 200
	Code int
}

type UserHeartsResponse struct {

	//Success Code, Ususally 200
	Code int

	//Account hearts
	Hearts string
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusBadRequest)
}

var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
}

var Messages []tools.CoinDetails

var MessagesLock sync.Mutex
