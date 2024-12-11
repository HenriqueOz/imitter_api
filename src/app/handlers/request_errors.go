package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RequestError struct {
	statusCode int
	err        error
}

var (
	ErrInternalServerError error = errors.New("internal server error")
)

func SendError(w http.ResponseWriter, requestError *RequestError) {
	w.WriteHeader(requestError.statusCode)

	response := map[string]string{
		"error": requestError.err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}
