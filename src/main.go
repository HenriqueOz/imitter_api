package main

import (
	"fmt"

	dotenv "github.com/joho/godotenv"
	"sm.com/m/src/app"
)

func main() {
	err := LoadEnvironment()
	if err != nil {
		fmt.Printf("error loading env file: %v\n", err)
	}

	app.Run()
}

// TODO move this function to a config.go file
func LoadEnvironment() (err error) {
	err = dotenv.Load("../.env")
	if err != nil {
		return
	}
	return
}
