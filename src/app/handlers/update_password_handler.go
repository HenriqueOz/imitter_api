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
	var requestBody UpdatePasswordRequest

	err := c.ShouldBindBodyWithJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	err = services.UpdateUserPassword(uuid, requestBody.NewPassword, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
