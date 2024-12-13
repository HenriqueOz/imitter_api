package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type SignInSuccessPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.UserSignIn

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Printf("error decoding body: %v", err)
		utils.SendError(w, &utils.RequestError{
			Message:    apperrors.ErrUnexpectedError.Error(),
			Err:        apperrors.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	if missing, err := validateSignInPayload(payload); err != nil {
		utils.SendErrorWithDetails(w, &utils.RequestError{
			Message:    apperrors.ErrMissingRequiredFields.Error(),
			Err:        apperrors.ErrBadRequest,
			StatusCode: http.StatusBadRequest,
			Details:    missing,
		})
	}

	if err := createUser(payload); err != nil {
		fmt.Printf("error creating user: %v", err)
		utils.SendError(w, &utils.RequestError{
			Err:        apperrors.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	utils.SendSuccess(w, SignInSuccessPayload{
		Name:  payload.Name,
		Email: payload.Email,
	}, http.StatusCreated)
}

func validateSignInPayload(payload models.UserSignIn) (map[string]any, error) {
	missing := utils.GetMissingFields([]string{"Name", "Password", "Email"}, payload)
	if len(missing) > 0 {
		return missing, apperrors.ErrMissingRequiredFields
	}
	return nil, nil
}

func createUser(model models.UserSignIn) error {
	return services.CreateUser(&models.UserSignIn{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	})
}
