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
		err := r.ParseForm()
		if err != nil {
			return NewHTTPError(err, 400, "")
		}
		dto := domain.CreateMessageDTO{
			From:    r.Form.Get("from"),
			To:      r.Form.Get("to"),
			Message: r.Form.Get("message"),
		}
		validationResult, _, err := registerService(&dto)
		if err != nil {
			return NewHTTPError(err, 400, "")
		}
		if validationResult != nil && validationResult.IsValid() {
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
