package router

import (
	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/handlers"
	"sm.com/m/src/app/middlewares"
)

type AppRouter struct{}

func NewAppRouter() *AppRouter {
	return &AppRouter{}
}

func (*AppRouter) BindAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/create", handlers.CreateAccountHandler)
		auth.POST("/refresh", middlewares.AuthMiddleware(), handlers.RefreshHandler)
		auth.POST("/logout", middlewares.AuthMiddleware(), handlers.LogoutHandler)
		auth.GET("/test", middlewares.AuthMiddleware(), handlers.AuthTestHandler)
	}
}

func (*AppRouter) BindUserRoutes(router *gin.RouterGroup) {
	user := router.Group("user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.PATCH("/update-name", handlers.UpdateNameHandler)
		user.PATCH("/update-password", handlers.UpdatePasswordHandler)
		user.DELETE("/delete", handlers.DeleteAccoutnHandler)
		// TODO Upload and Download user avatar
	}
}

func (*AppRouter) BindPostRoutes(router *gin.RouterGroup) {
	posts := router.Group("/posts")
	posts.Use(middlewares.AuthMiddleware())
	{
		posts.POST("/create", handlers.CreatePostHandler)
		posts.GET("/recent", handlers.RecentPostsHandler)
	}
	// GET following posts
	// DELETE delete post
}

func (*AppRouter) BindSearchRoutes(router *gin.RouterGroup) {}
