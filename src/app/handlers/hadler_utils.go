package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendError(w http.ResponseWriter, requestError *RequestError) {
	w.WriteHeader(requestError.StatusCode)

	response := map[string]string{
		"error":   requestError.Err.Error(),
		"message": requestError.Message,
	}

	json.NewEncoder(w).Encode(response)

	w.WriteHeader(requestError.StatusCode)
}

func SendErrorWithDetails(w http.ResponseWriter, requestError *RequestError) {
	response := map[string]any{
		"error":   requestError.Err.Error(),
		"message": requestError.Message,
	}

	if requestError.Details != nil {
		response["details"] = requestError.Details
	}

	json.NewEncoder(w).Encode(response)
}

func ValidateRequiredFields(w http.ResponseWriter, body io.ReadCloser, requiredFields []string) bool {
	var payload map[string]any = make(map[string]any)

	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		fmt.Printf("error decoding body: %v", err)
		return false
	}

	missing := GetRequiredFields(requiredFields, payload)
	return len(missing) > 0
}

func GetRequiredFields(requiredFields []string, input map[string]any) (missing map[string]any) {
	missing = make(map[string]any)

	for _, field := range requiredFields {
		if _, ok := input[field]; !ok {
			missing[field] = "required"
		}
	}
	return missing
}
