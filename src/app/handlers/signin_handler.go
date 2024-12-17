package handlers

import (
	"encoding/json"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(`{ message: "auth done" }`)

	w.Write(json)
}
