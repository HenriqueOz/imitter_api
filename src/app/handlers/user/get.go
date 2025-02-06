package user

import "github.com/gin-gonic/gin"

func GetUserProfileByUUIDHandler(c *gin.Context) {
	uuid := c.Param("uuid")
}

type GetUserProfileByNameRequest struct {
	Query string `form:"q" binding:"required"`
}

func GetUserProfileByNameHandler(c *gin.Context) {
	//uuid := c.Param("uuid")
}
