package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/server/api"
	"github.com/RYANCOAL9999/DisSystem/server/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetUserHearts(writer http.ResponseWriter, request *http.Request) {

	var params = api.Params{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, request.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var UserHearts = (*database).GetUserHearts(params.Username)
	if UserHearts == nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var response = api.UserHeartsResponse{
		Hearts: (*UserHearts).Heart,
		Code:   http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

}
