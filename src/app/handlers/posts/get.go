package posts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type recentPostsRequest struct {
	ate string `form:"date" binding:"required"`
}

func MyRecentPostsHandler(c *gin.Context) {
	startDate, ok := getRecentStartTime(c)
	if !ok {
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetMyRecent(startDate, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsByUUIDHandler(c *gin.Context) {
	startDate, ok := getRecentStartTime(c)
	if !ok {
		return
	}

	uuid := c.GetHeader("uuid")
	postUserUUID := c.Param("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecentByPostUserUUID(startDate, uuid, postUserUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsHandler(c *gin.Context) {

	startDate, ok := getRecentStartTime(c)
	if !ok {
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecent(startDate, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsFollowingHandler(c *gin.Context) {

	startDate, ok := getRecentStartTime(c)
	if !ok {
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecentFollowing(startDate, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func getRecentStartTime(c *gin.Context) (startDate time.Time, ok bool) {
	request := recentPostsRequest{}

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return time.Time{}, false
	}

	startDate, err = time.Parse(time.DateTime, c.Query("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return time.Time{}, false
	}

	return startDate, true
}
