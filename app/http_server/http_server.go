package httpserver

import (
	"net"
	"net/http"
	"os"
	"time"

	"sm.com/m/app/handlers"
	"sm.com/m/app/middlewares"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	mux = assignRoutes(mux)

	var handler http.Handler = mux
	handler = assignMiddlwares(handler)

	var server *http.Server = assignServer(handler)

	return server
}

func assignRoutes(mux *http.ServeMux) *http.ServeMux {
	// TODO do the routes assigning stuff
	mux.HandleFunc("/", handlers.HomeHandler)
	return mux
}

func assignMiddlwares(handler http.Handler) http.Handler {
	// TODO do the middleware assigning stuff
	handler = middlewares.RequestLogger(handler)
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
