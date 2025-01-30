package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type LikeRequest struct {
	PostId uint64 `json:"post_id" binding:"required"`
}

func LikeHandler(c *gin.Context) {
	request := LikeRequest{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	service := services.NewPostService()
	err = service.ToogleLike(c.GetHeader("uuid"), request.PostId)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
