package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if !ValidateRequiredFields(w, r.Body, []string{
		"name",
		"email",
		"password",
	}) {
		return
	}

	var model models.UserSignIn
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		fmt.Printf("error decoding body: %v", err)
		SendError(w, &RequestError{
			Err:        ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
			Message:    ErrUnexpectedError.Error(),
		})
		return
	}

	err = createUser(model)
	if err != nil {
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
