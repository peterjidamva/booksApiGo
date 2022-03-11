package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "drift"
	hostname = "localhost"
	dbname   = "library"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func init() {
	sql.Register("mysql", &MySQLDriver{})
}
