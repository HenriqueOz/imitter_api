package router

import (
	"github.com/gin-gonic/gin"

	"sm.com/m/src/app/handlers/auth"
	"sm.com/m/src/app/handlers/posts"
	"sm.com/m/src/app/handlers/user"
	"sm.com/m/src/app/middlewares"
)

type AppRouter struct{}

func NewAppRouter() *AppRouter {
	return &AppRouter{}
}

func (*AppRouter) BindAuthRoutes(router *gin.RouterGroup) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", auth.LoginHandler)
		authRoutes.POST("/create", auth.CreateAccountHandler)
		authRoutes.POST("/refresh", middlewares.AuthMiddleware(), auth.RefreshHandler)
		authRoutes.POST("/logout", middlewares.AuthMiddleware(), auth.LogoutHandler)
		authRoutes.GET("/test", middlewares.AuthMiddleware(), auth.AuthTestHandler)
	}
}

func (*AppRouter) BindUserRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.PATCH("/update-name", user.UpdateNameHandler)
		userRoutes.PATCH("/update-password", user.UpdatePasswordHandler)
		userRoutes.DELETE("/delete", user.DeleteAccoutnHandler)
		userRoutes.POST("/follow", user.ToogleFollowHandler)
		userRoutes.GET("/:uuid/profile", user.GetUserProfileByUUIDHandler)
	}
}

func (*AppRouter) BindPostRoutes(router *gin.RouterGroup) {
	postsRoutes := router.Group("/posts")
	postsRoutes.Use(middlewares.AuthMiddleware())
	{
		postsRoutes.POST("/create", posts.CreatePostHandler)
		postsRoutes.GET("/recent", posts.RecentPostsHandler)
		postsRoutes.GET("/recent/me", posts.MyRecentPostsHandler)
		postsRoutes.GET("/:uuid/recent", posts.RecentPostsByUUIDHandler)
		postsRoutes.POST("/like", posts.ToogleLikeHandler)
		postsRoutes.GET("/following", posts.RecentPostsFollowingHandler)
		postsRoutes.DELETE("/delete", posts.DeletePostHandler)
	}
}

func (*AppRouter) BindSearchRoutes(router *gin.RouterGroup) {
	searchRoutes := router.Group("search")
	searchRoutes.Use(middlewares.AuthMiddleware())
	{
	}
}
