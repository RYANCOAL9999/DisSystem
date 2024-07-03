package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/server/api"
	log "github.com/sirupsen/logrus"
)

func PollingHandler(writer http.ResponseWriter, request *http.Request) {

	api.MessagesLock.Lock()
	defer api.MessagesLock.Unlock()

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(api.Messages)
	if err != nil {
		log.Printf("Error encoding messages: %v\n", err)
		http.Error(writer, "Failed to encode messages", http.StatusInternalServerError)
		return
	}

	// Clear the messages after sending them to the client
	api.Messages = nil
}
