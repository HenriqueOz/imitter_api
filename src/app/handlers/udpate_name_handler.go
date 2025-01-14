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
	Name     string `json:"name" binding:"required"`
}

func UpdateNameHandler(c *gin.Context) {
	userService := services.NewUserService()
	var requestBody UpdateNameRequest

	err := c.ShouldBindBodyWithJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	err = userService.UpdateUserName(uuid, requestBody.Name, requestBody.NewName, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
