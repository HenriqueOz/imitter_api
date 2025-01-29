package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type LoginRequest struct {
	Method   string `json:"method" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var err error
	service := services.NewAuthService()
	var requestBody LoginRequest

	err = c.ShouldBindJSON(&requestBody)

	if err != nil {
		log.Printf("%v\n", err)
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

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
