package middleware

import (
	"errors"
	"net/http"

	"github.com/RYANCOAL9999/DisSystem/server/api"
	"github.com/RYANCOAL9999/DisSystem/server/internal/tools"
	log "github.com/sirupsen/logrus"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		errUnAuthorized := errors.New("unauthorized access")

		if username == "" {
			api.RequestErrorHandler(w, errUnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(errUnAuthorized)
			api.RequestErrorHandler(w, errUnAuthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
