package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/utils"
)

func RefreshHandler(c *gin.Context) {
	uuid := c.GetHeader("uuid")

	payload, err := GetTokenPayload(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(
			err,
			nil,
		))
		return
	}

	c.JSON(http.StatusNoContent, payload)
}

func GetTokenPayload(uuid string) (map[string]any, error) {
	accessToken, err := utils.GenerateJwtToken(uuid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshJwtToken(accessToken)
	if err != nil {
		return nil, err
	}

	payload := map[string]any{
		"access_token":  "Bearer " + accessToken,
		"refresh_token": "Bearer " + refreshToken,
	}

	return payload, nil
}
