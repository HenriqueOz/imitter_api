package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type UpdateNameRequest struct {
	Password string `json:"password" binding:"required"`
	NewName  string `json:"new_name" binding:"required"`
}

func UpdateNameHandler(c *gin.Context) {
	var requestBody UpdateNameRequest
	err := c.ShouldBindBodyWithJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewUserService()
	err = service.UpdateUserName(uuid, requestBody.NewName, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
