package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/utils"
)

func RefreshHandler(c *gin.Context) {
	uuid := c.GetHeader("uuid")

	tokenPayload, err := getTokenPayload(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(
			apperrors.ErrUnexpected,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(tokenPayload))
}

func getTokenPayload(uuid string) (map[string]any, error) {
	accessToken, err := utils.GenerateJwtToken(uuid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshJwtToken(uuid, accessToken)
	if err != nil {
		return nil, err
	}

	payload := map[string]any{
		"access_token":  "Bearer " + accessToken,
		"refresh_token": "Bearer " + refreshToken,
	}

	return payload, nil
}
