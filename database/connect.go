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

var db *sql.DB

type Book struct {
	Id    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`

	//taking a reference of Author from it's given structure
	Author *Author `json:"author"`
}
type Author struct {
	authorid  string `json:"authid"`
	FirstName string `json:"first_name"`
	lastname  string `json:"last_name"`
	age       string `json:"age"`
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func getProduct() (err error) {

	db, err = sql.Open("mysql", "root:drift@/library")

	var p Book
	err = db.QueryRow("SELECT title FROM books").Scan(&p.Title)

	if err != nil {
		fmt.Println(err.Error())
	}

  	fmt.Println(p.Title)
    
	  return err
}

func main() {
   getProduct();

}