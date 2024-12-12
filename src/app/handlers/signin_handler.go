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

	if missing := getSignInPayloadMissingFields(payload); len(missing) > 0 {
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

func getSignInPayloadMissingFields(payload models.UserSignIn) (missing map[string]string) {
	missing = make(map[string]string)

	if payload.Email == "" {
		missing["email"] = "required"
	}
	if payload.Name == "" {
		missing["name"] = "required"
	}
	if payload.Password == "" {
		missing["password"] = "required"
	}

	return missing
}

func createUser(model models.UserSignIn) error {
	return repositories.CreateUser(models.UserSignIn{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	})
}
