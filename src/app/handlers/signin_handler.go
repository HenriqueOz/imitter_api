package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]any = map[string]any{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Printf("error decoding body: %v", err)
		SendError(w, &RequestError{
			err:        ErrInternalServerError,
			statusCode: http.StatusInternalServerError,
		})
		return
	}

	missing := getMissingFields(payload)
	if len(missing) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(missing)
		return
	}

	err = createUser(payload)
	if err != nil {
		fmt.Printf("error creating user: %v", err)
		SendError(w, &RequestError{
			err:        ErrInternalServerError,
			statusCode: http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "user created",
	})
}

func getMissingFields(payload map[string]any) map[string]string {
	return utils.VerifyRequiredFields([]string{
		"name", "email", "password",
	}, payload)
}

func createUser(payload map[string]any) error {
	return repositories.CreateUser(models.UserSignIn{
		Name:     payload["name"].(string),
		Email:    payload["email"].(string),
		Password: payload["password"].(string),
	})
}
