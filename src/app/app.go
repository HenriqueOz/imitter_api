package app

import (
	httpserver "sm.com/m/src/app/http_server"
)

func Run() {
	server := httpserver.NewServer()
	server.Run()
}
