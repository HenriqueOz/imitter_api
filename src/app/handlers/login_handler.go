package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type LoginRequestBody struct {
	Method   string `json:"method" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var err error
	var requestBody LoginRequestBody

	err = c.ShouldBindJSON(&requestBody)

	if err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, utils.ResponseError(
				apperrors.ErrMissingFields,
				utils.DescriptiveError(verr),
			))
		} else {
			c.JSON(http.StatusBadRequest, utils.ResponseError(
				apperrors.ErrInvalidRequest,
				err.Error(),
			))
		}
		return
	}

	responseData, err := services.Login(requestBody.Method, requestBody.Login, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrLogin,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(responseData))
}
