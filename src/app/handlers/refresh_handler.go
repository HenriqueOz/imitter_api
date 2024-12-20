package handlers

import (
	"encoding/json"
	"net/http"

	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	payload := generateNewTokens(w, &models.UserSignIn{Uuid: r.Header["Uuid"][0]})
	if payload == nil {
		return
	}

	json.NewEncoder(w).Encode(payload)
}

func generateNewTokens(w http.ResponseWriter, user *models.UserSignIn) *models.UserAuth {
	accessToken, err := utils.GenerateJwtToken(user)
	if err != nil {
		utils.SendInternalServerError(w)
		return nil
	}

	refreshToken, err := utils.GenerateRefreshJwtToken(accessToken)
	if err != nil {
		utils.SendInternalServerError(w)
		return nil
	}

	return &models.UserAuth{
		AccessToken:  "Bearer " + accessToken,
		RefreshToken: "Bearer " + refreshToken,
	}
}
