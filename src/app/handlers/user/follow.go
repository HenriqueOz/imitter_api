package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type toogleFollowRequest struct {
	ToFollowUUID string `json:"to_follow_uuid" binding:"required"`
}

func ToogleFollowHandler(c *gin.Context) {
	var request toogleFollowRequest
	err := c.ShouldBindBodyWithJSON(&request)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewFollowService()
	err = service.ToogleFollow(uuid, request.ToFollowUUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
