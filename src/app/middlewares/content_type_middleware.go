package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
