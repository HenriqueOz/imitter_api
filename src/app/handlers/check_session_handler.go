package handlers

import (
	"net/http"
)

func CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)

	{
		var a int = 1
	}
}
