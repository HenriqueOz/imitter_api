package database

import (
	"database/sql"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	connectionConfig *mysql.Config
	Conn             *sql.DB
)

func initConfig() {
	if connectionConfig == nil {
		connectionConfig = &mysql.Config{
			User:                 os.Getenv("DB_USER"),
			Passwd:               os.Getenv("DB_PASSWORD"),
			Addr:                 net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			DBName:               os.Getenv("DB_NAME"),
			Net:                  "tcp",
			AllowNativePasswords: true,
		}
	}
}

func OpenConnection() error {
	initConfig()
	dsn := connectionConfig.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	Conn = db
	return nil
}

func CloseConnection() error {
	if Conn == nil {
		return nil
	}

	err := Conn.Close()
	if err != nil {
		return err
	}

	return nil
}
