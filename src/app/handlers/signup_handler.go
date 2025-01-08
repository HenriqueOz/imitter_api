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

type SignUpSuccessPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.UserSignUp

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Printf("error decoding body: %v\n", err)
		utils.SendInternalServerError(w)
		return
	}

	if missing, err := validateSignUpPayload(payload); err != nil {
		utils.SendMissingFieldsError(w, missing)
		return
	}

	if err := createUser(payload); err != nil {
		fmt.Printf("error creating user: %v\n", err)
		sendCreateUserError(w, err)
		return
	}

	utils.SendSuccess(w, SignUpSuccessPayload{
		Name:  payload.Name,
		Email: payload.Email,
	}, 201)
}

func validateSignUpPayload(payload models.UserSignUp) (map[string]any, error) {
	missing, _ := utils.GetMissingFields([]string{"Name", "Password", "Email"}, payload)

	if len(missing) > 0 {
		return missing, apperrors.ErrMissingRequiredFields
	}

	return nil, nil
}

func sendCreateUserError(w http.ResponseWriter, err error) {
	utils.SendError(w, &utils.RequestError{
		Err:        apperrors.ErrCreateUser,
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	})
}

func createUser(model models.UserSignUp) error {
	return services.CreateUser(&models.UserSignUp{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	})
}
