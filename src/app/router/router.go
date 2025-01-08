package router

import (
	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/handlers"
	"sm.com/m/src/app/middlewares"
)

func BindAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/create-account", handlers.CreateAccountHandler)
		auth.GET("/refresh", middlewares.AuthMiddleware(), handlers.RefreshHandler)
		auth.GET("/check-auth", middlewares.AuthMiddleware(), handlers.CheckAuthHandler)
	}
}
