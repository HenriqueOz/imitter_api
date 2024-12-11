package httpserver

import (
	"net"
	"net/http"
	"os"
	"time"

	"sm.com/m/src/app/handlers"
	"sm.com/m/src/app/middlewares"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux = assignRoutes(mux)
	handler := assignMiddlwares(mux)

	return assignServer(handler)
}

func assignRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("POST /signin", handlers.SignInHandler)
	return mux
}

func assignMiddlwares(handler http.Handler) http.Handler {
	return chainMiddlewares(
		handler,
		middlewares.RequestLoggerMiddleware,
		middlewares.ContentTypeMiddleware,
		middlewares.CorsMiddleware,
	)
}

func chainMiddlewares(handler http.Handler, middlewares ...middlewares.Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func assignServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         net.JoinHostPort(os.Getenv("SVHOST"), os.Getenv("SVPORT")),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
