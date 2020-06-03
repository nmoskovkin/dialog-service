package controller

import (
	"database/sql"
	"dialogService/app/repository"
	"dialogService/domain"
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Status string
}

func CreateMessagePostHandler(db *sql.DB) ErrorReturningHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		registerService := domain.CreateMessageService(repository.CreateMessageWriteRepository(db))
		var dto domain.CreateMessageDTO

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			return NewHTTPError(err, 400, "")
		}
		validationResult, _, err := registerService(&dto)
		if err != nil {
			return NewHTTPError(err, 400, "")
		}
		if validationResult != nil && !validationResult.IsValid() {
			return NewHTTPError(err, 400, "")
		}
		js, err := json.Marshal(MessageResponse{Status: "OK"})
		if err != nil {
			return NewHTTPError(err, 500, "")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		return nil
	}
}
