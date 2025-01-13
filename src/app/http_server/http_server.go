package httpserver

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/middlewares"
	"sm.com/m/src/app/router"
)

type Server struct {
	Port string
	Host string
	Gin  *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		Port: os.Getenv("SERVER_PORT"),
		Host: os.Getenv("SERVER_HOST"),
		Gin:  gin.New(),
	}

	server.Gin.SetTrustedProxies(nil) //! Not safe

	server.setMiddlewares()
	server.setRoutes()

	return server
}

func (server *Server) Run() {
	address := net.JoinHostPort(server.Host, server.Port)
	err := server.Gin.Run(address)
	if err != nil {
		log.Printf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func (server *Server) setMiddlewares() {
	server.Gin.Use(gin.Logger())
	server.Gin.Use(gin.Recovery())
	server.Gin.Use(middlewares.ContentTypeMiddleware())
	server.Gin.Use(middlewares.CorsMiddleware())
}

func (server *Server) setRoutes() {
	v1 := server.Gin.Group("/v1")
	router.BindAuthRoutes(v1)
	router.BindUserRoutes(v1)
}
