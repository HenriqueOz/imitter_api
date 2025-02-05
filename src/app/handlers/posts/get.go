package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

type recentPostsRequest struct {
	Limit  int `form:"limit" binding:"required"`
	Offset int `form:"offset"`
}

func MyRecentPostsHandler(c *gin.Context) {
	request := recentPostsRequest{}

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")
	service := services.NewPostService()

	posts, err := service.GetMyRecent(request.Limit, request.Offset, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsByUUIDHandler(c *gin.Context) {
	request := recentPostsRequest{}

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")
	postUserUUID := c.Param("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecentByPostUserUUID(request.Limit, request.Offset, uuid, postUserUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsHandler(c *gin.Context) {
	request := recentPostsRequest{}

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecent(request.Limit, request.Offset, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}

func RecentPostsFollowingHandler(c *gin.Context) {
	request := recentPostsRequest{}

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	uuid := c.GetHeader("uuid")

	service := services.NewPostService()
	posts, err := service.GetRecentFollowing(request.Limit, request.Offset, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(apperrors.ErrInvalidRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(posts))
}
