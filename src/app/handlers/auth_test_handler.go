package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/utils"
)

func AuthTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseSuccess(map[string]interface{}{
		"status": "you're authenticated!",
	}))
}
