package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type createPostRequest struct {
	Content string `json:"content" binding:"required"`
}

func CreatePostHandler(c *gin.Context) {
	var err error
	var request createPostRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	service := services.NewPostService()
	err = service.CreatePost(c.GetHeader("uuid"), request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.Render(http.StatusCreated, render.Data{})
}
