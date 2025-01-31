package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type loginRequest struct {
	Method   string `json:"method" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var err error
	var requestBody loginRequest

	err = c.ShouldBindJSON(&requestBody)

	if err != nil {
		log.Printf("%v\n", err)
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	service := services.NewAuthService()
	responseData, err := service.Login(requestBody.Method, requestBody.Login, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrLogin,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(responseData))
}

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

func AuthTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseSuccess(map[string]interface{}{
		"status": "you're authenticated!",
		"uuid":   c.GetHeader("uuid"),
	}))
}
