package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/services"
	"sm.com/m/src/app/utils"
)

func GetUserProfileByUUIDHandler(c *gin.Context) {
	service := services.NewUserService()
	result, err := service.GetUserProfileByUUID(c.GetHeader("uuid"), c.Param("uuid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(result))
}

type GetUserProfileByNameRequest struct {
	Query string `form:"q" binding:"required"`
}

func GetUserProfileByNameHandler(c *gin.Context) {
	var request GetUserProfileByNameRequest
	err := c.ShouldBindQuery(&request)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}

	service := services.NewUserService()
	result, err := service.GetUserProfileByName(c.GetHeader("uuid"), request.Query)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(result))
}
