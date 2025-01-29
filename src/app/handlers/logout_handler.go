package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func LogoutHandler(c *gin.Context) {
	var err error

	tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]
	token, err := utils.ParseToken(tokenString)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidToken,
			apperrors.ErrInvalidClaims.Error(),
		))
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	jti, ok := claims["jti"]
	if !ok {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidToken,
			apperrors.ErrTokenFormat.Error(),
		))
		return
	}

	blackListService := services.NewBlackListService()
	err = blackListService.AddTokenToBlacklist(jti.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidToken,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
