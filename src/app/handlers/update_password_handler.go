package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type UpdatePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func UpdatePasswordHandler(c *gin.Context) {
	service := services.NewUserService()

	var requestBody UpdatePasswordRequest

	err := c.ShouldBindBodyWithJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	err = service.UpdateUserPassword(uuid, requestBody.NewPassword, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
