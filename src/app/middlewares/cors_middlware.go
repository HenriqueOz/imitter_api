package middlewares

import (
	"github.com/gin-gonic/gin"
)

// TODO not in use middleware
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Next()
	}
}
