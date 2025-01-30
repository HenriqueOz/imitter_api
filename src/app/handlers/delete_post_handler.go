package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type DeletePostRequest struct {
	PostId       uint64 `json:"id" binding:"required"`
	PostUserUUID string `json:"user_uuid" binding:"required"`
}

func DeletePostHandler(c *gin.Context) {
	var err error
	var request DeletePostRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	service := services.NewPostService()
	err = service.DeletePost(request.PostUserUUID, c.GetHeader("uuid"), request.PostId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
