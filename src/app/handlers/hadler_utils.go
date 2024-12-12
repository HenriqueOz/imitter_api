package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

func SendError(w http.ResponseWriter, requestError *RequestError) {
	w.WriteHeader(requestError.StatusCode)

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

func GetMissingFields(requiredFields []string, model interface{}) (missing map[string]any) {
	missing = make(map[string]any)
	mappedStruct := structs.Map(model)

	for _, field := range requiredFields {
		value := reflect.ValueOf(mappedStruct[field])

		if value.IsZero() {
			missing[strings.ToLower(field)] = "required"
		}
	}

	return missing
}
