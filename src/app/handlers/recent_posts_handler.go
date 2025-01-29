package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func RecentPostsHandler(c *gin.Context) {
	startDate, err := time.Parse(time.DateTime, c.Query("start_datetime"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	service := services.NewPostService()
	posts, err := service.GetRecentByStartDate(startDate, c.GetHeader("uuid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}
