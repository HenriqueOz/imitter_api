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
		userRoutes.PATCH("/name", user.UpdateNameHandler)
		userRoutes.PATCH("/password", user.UpdatePasswordHandler)
		userRoutes.DELETE("/", user.DeleteAccoutnHandler)
		userRoutes.POST("/follow", user.ToogleFollowHandler)
		userRoutes.GET("/:uuid", user.GetUserProfileByUUIDHandler)
		userRoutes.GET("/search/", user.GetUserProfileByNameHandler)
	}
}

func (*AppRouter) BindPostRoutes(router *gin.RouterGroup) {
	postsRoutes := router.Group("/posts")
	postsRoutes.Use(middlewares.AuthMiddleware())
	{
		postsRoutes.POST("/", posts.CreatePostHandler)
		postsRoutes.POST("/like", posts.ToogleLikeHandler)

		postsRoutes.GET("/", posts.RecentPostsHandler)
		postsRoutes.GET("/me", posts.MyRecentPostsHandler)
		postsRoutes.GET("/:uuid", posts.RecentPostsByUUIDHandler)
		postsRoutes.GET("/following", posts.RecentPostsFollowingHandler)

		postsRoutes.DELETE("/", posts.DeletePostHandler)
	}
}

func (*AppRouter) BindSearchRoutes(router *gin.RouterGroup) {
	searchRoutes := router.Group("search")
	searchRoutes.Use(middlewares.AuthMiddleware())
	{
	}
}
