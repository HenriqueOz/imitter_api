package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"frango": "frito",
	}

	json.NewEncoder(w).Encode(response)
}
