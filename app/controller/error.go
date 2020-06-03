package controller

import (
	"encoding/json"
	"net/http"
)

type ClientError interface {
	Error() string
}

type HTTPError struct {
	Cause  error
	Detail string
	Status int
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

func CreateErrorHandlerFunc() ErrorHandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request, clientError ClientError) {
		js, err := json.Marshal(clientError)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(js)
	}
}
