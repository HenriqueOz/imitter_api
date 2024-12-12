package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.UserSignIn

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Printf("error decoding body: %v", err)
		SendError(w, &RequestError{
			Message:    ErrUnexpectedError.Error(),
			Err:        ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	if missing := GetMissingFields([]string{"name", "password", "email"}, payload); len(missing) > 0 {
		SendErrorWithDetails(w, &RequestError{
			Message:    ErrMissingRequiredFields.Error(),
			Err:        ErrBadRequest,
			StatusCode: http.StatusBadRequest,
			Details:    missing,
		})
		return
	}

	if err := createUser(payload); err != nil {
		fmt.Printf("error creating user: %v", err)
		SendError(w, &RequestError{
			Err:        ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
			Message:    ErrUnexpectedError.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "user created",
	})
}

func createUser(model models.UserSignIn) error {
	return repositories.CreateUser(models.UserSignIn{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	})
}
