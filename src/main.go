package main

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	"sm.com/m/src/app"
	db "sm.com/m/src/app/database"
)

func main() {
	defer CloseDB()

	var err error

	err = LoadEnvironment()
	if err != nil {
		log.Printf("Failed to load .env file: %v\n", err)
		os.Exit(1)
	}

	err = db.OpenConnection()
	if err != nil {
		log.Printf("Failed to open database connection: %v\n", err)
		os.Exit(1)
	}

	app.Run()
}

func LoadEnvironment() (err error) {
	err = dotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	err := db.CloseConnection()
	if err != nil {
		log.Printf("Failed to close database connection: %v\n", err)
	}
}
