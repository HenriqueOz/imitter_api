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
			User:                 os.Getenv("DBUSER"),
			Passwd:               os.Getenv("DBPASSWORD"),
			Addr:                 net.JoinHostPort(os.Getenv("DBHOST"), os.Getenv("DBPORT")),
			DBName:               os.Getenv("DBNAME"),
			Net:                  "tcp",
			AllowNativePasswords: true,
		}
	}
}

func OpenConnection() (err error) {
	initConfig()
	dsn := connectionConfig.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	Conn = db
	return
}

func CloseConnection() (err error) {
	if Conn != nil {
		err = Conn.Close()
		if err != nil {
			return
		}
	}
	return
}
