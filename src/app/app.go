package app

import (
	"fmt"

	db "sm.com/m/src/app/database"
	httpserver "sm.com/m/src/app/http_server"
)

func Run() {
	defer cleanUp()

	if err := db.OpenConnection(); err != nil {
		fmt.Printf("error opnening database connection: %v\n", err)
	}

	server := httpserver.NewServer()

	fmt.Printf("listening and serving at http://%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("error listening and serving: %v\n", err)
	}

}

func cleanUp() {
	if err := db.CloseConnection(); err != nil {
		fmt.Printf("error closing database connection: %v\n", err)
	}
}
