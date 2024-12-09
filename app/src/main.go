package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/sm_database")

	if err != nil {
		panic(err)
	}

	return db, fmt.Errorf("error while opening database connection")
}

func main() {
	db, err := connection()

	if err != nil {
		result, err := db.Query("SELECT * FROM user")
		if err != nil {

		}

		fmt.Println(result.Columns())
	}
}
