package api

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Params
type Params struct {
	Username string
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
