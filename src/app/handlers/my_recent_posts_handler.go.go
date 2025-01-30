package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func MyRecentPostsHandler(c *gin.Context) {
	lastPostId, err := strconv.Atoi(c.Query("last_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetMyRecent(uuid, uint64(lastPostId))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}
