package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
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
	service := services.NewAuthService()

	var err error
	var requestBody CreateAccountRequestBody

	err = c.ShouldBindJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	err = service.CreateUser(&models.UserModel{
		Email:    requestBody.Email,
		Name:     requestBody.Name,
		Password: requestBody.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.Render(http.StatusCreated, render.Data{})
}
