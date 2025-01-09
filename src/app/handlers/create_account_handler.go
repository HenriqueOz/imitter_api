package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type CreateAccountRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateAccountHandler(c *gin.Context) {
	var err error
	var requestBody CreateAccountRequestBody

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

	err = services.CreateUser(&models.UserModel{
		Email:    requestBody.Email,
		Name:     requestBody.Name,
		Password: requestBody.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
	}

	c.Render(http.StatusCreated, render.Data{})
}
