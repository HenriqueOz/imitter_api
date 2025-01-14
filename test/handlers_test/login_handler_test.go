package handlerstest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/handlers"
	httpserver "sm.com/m/src/app/http_server"
)

type LoginRequest handlers.LoginRequestBody

type DetailsMissinField struct {
	Field   string
	Message string
}

type LoginError struct {
	Message string            `json:"error"`
	Details map[string]string `json:"details"`
}

type LoginMissingFieldsInput struct {
	Request LoginRequest
	Error   LoginError
}

func TestLoginMissingFields(t *testing.T) {
	server := httpserver.NewServer()
	inputs := []LoginMissingFieldsInput{
		{
			Request: LoginRequest{Method: "name", Login: "random@emial"},
			Error: LoginError{
				Message: apperrors.ErrMissingFields.Error(),
				Details: map[string]string{
					"Password": "required, string",
				},
			},
		},
		{
			Request: LoginRequest{Method: "name"},
			Error: LoginError{
				Message: apperrors.ErrMissingFields.Error(),
				Details: map[string]string{
					"Password": "required, string",
					"Login":    "required, string",
				},
			},
		},
		{
			Request: LoginRequest{},
			Error: LoginError{
				Message: apperrors.ErrMissingFields.Error(),
				Details: map[string]string{
					"Method":   "required, string",
					"Password": "required, string",
					"Login":    "required, string",
				},
			},
		},
	}

	for _, input := range inputs {
		w := httptest.NewRecorder()
		inputJson, _ := json.Marshal(input.Request)

		request, _ := http.NewRequest("POST", "/v1/auth/login", strings.NewReader(string(inputJson)))
		server.Gin.ServeHTTP(w, request)

		response := LoginError{}
		json.Unmarshal([]byte(w.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, input.Error.Details, response.Details)
		assert.Equal(t, input.Error.Message, response.Message)
	}
}
