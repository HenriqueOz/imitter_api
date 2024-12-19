package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	apperrors "sm.com/m/src/app/app_errors"
)

type RequestError struct {
	StatusCode int
	Err        error
	Message    string
	Details    map[string]any
}

func SendError(w http.ResponseWriter, requestError *RequestError) {
	response := map[string]string{
		"error":   requestError.Err.Error(),
		"message": requestError.Message,
	}

	w.WriteHeader(requestError.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func SendErrorWithDetails(w http.ResponseWriter, requestError *RequestError) {
	response := map[string]any{
		"error":   requestError.Err.Error(),
		"message": requestError.Message,
	}

	if requestError.Details != nil {
		response["details"] = requestError.Details
	}

	w.WriteHeader(requestError.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func SendSuccess(w http.ResponseWriter, payload interface{}, status int) error {
	if !structs.IsStruct(payload) {
		return apperrors.ErrIsNotStruct
	}

	if status == 0 {
		status = http.StatusOK
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
	return nil
}

func GetMissingFields(requiredFields []string, data interface{}) (missing map[string]any, err error) {
	if !structs.IsStruct(data) {
		return nil, apperrors.ErrIsNotStruct
	}

	missing = make(map[string]any)
	mappedStruct := structs.Map(data)

	for _, field := range requiredFields {
		value := reflect.ValueOf(mappedStruct[field])
		if !value.IsValid() || value.IsZero() {
			missing[strings.ToLower(field)] = "required"
		}
	}

	return missing, nil
}

func SendMissingFieldsError(w http.ResponseWriter, missing map[string]any) {
	SendErrorWithDetails(w, &RequestError{
		StatusCode: 400,
		Err:        apperrors.ErrBadRequest,
		Message:    apperrors.ErrMissingRequiredFields.Error(),
		Details:    missing,
	})
}

func SendInternalServerError(w http.ResponseWriter) {
	SendError(w, &RequestError{
		StatusCode: 500,
		Err:        apperrors.ErrInternalServerError,
		Message:    apperrors.ErrUnexpectedError.Error(),
	})
}
