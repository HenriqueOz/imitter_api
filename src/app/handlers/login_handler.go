package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type LoginRequestBody struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var err error
	var requestBody LoginRequestBody

	err = c.Bind(&requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	responseBody, err := services.Login(requestBody.Login, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(
			apperrors.ErrLogin,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, responseBody)
}
