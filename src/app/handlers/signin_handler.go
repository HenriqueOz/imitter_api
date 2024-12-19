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

type SignInPayload struct {
	Email    string `json:"emaiL"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInSuccessPayload struct {
	AccessToken  string `jsom:"access_token"`
	RefreshToken string `jsom:"refresh_token"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload SignInPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Printf("error decoding body: %v\n", err)
		utils.SendInternalServerError(w)
		return
	}

	if missing, err := validateSignInPayload(payload); err != nil {
		utils.SendMissingFieldsError(w, missing)
		return
	}

	var err error
	var userAuth *models.UserAuth
	if payload.Email != "" {
		userAuth, err = services.SignInWithEmail(payload.Email, payload.Password)
	} else {
		userAuth, err = services.SignInWithName(payload.Name, payload.Password)
	}

	if err != nil {
		sendSignInError(w, err)
		return
	}

	utils.SendSuccess(w, userAuth, 200)
}

func sendSignInError(w http.ResponseWriter, err error) {
	utils.SendError(w, &utils.RequestError{
		Err:        apperrors.ErrInternalServerError,
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	})
}

func validateSignInPayload(payload interface{}) (map[string]any, error) {
	missing, _ := utils.GetMissingFields([]string{"Name", "Password", "Email"}, payload)

	_, emailExists := missing["email"]
	_, nameExists := missing["name"]

	if len(missing) == 1 {
		if !emailExists && !nameExists {
			return missing, apperrors.ErrMissingRequiredFields
		}
	} else if len(missing) > 1 {
		return missing, apperrors.ErrMissingRequiredFields
	}

	return nil, nil
}
