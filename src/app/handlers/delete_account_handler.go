package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type DeleteAccountRequest struct {
	Password string `json:"password" biding:"required"`
}

func DeleteAccoutnHandler(c *gin.Context) {
	requestBody := new(DeleteAccountRequest)

	err := c.ShouldBindBodyWithJSON(&requestBody)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewUserService()
	err = service.DeleteUserAccount(uuid, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
