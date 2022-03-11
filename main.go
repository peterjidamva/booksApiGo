package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//creating a structure of books

type Book struct {
	Id    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`

	//taking a reference of Author from it's given structure
	Author *Author `json:"author"`
}

//cretaing a structure of  author
type Author struct {
	authorid  string `json:"authid"`
	FirstName string `json:"first_name"`
	lastname  string `json:"last_name"`
	age       string `json:"age"`
}

//handlers functions

func getListOfBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//initiate a router
	r := mux.NewRouter()

	//route handling

	r.HandleFunc("/api/books", getListOfBooks).Methods("GET")
	r.HandleFunc("/api/books", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("GET")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	//starting server
	fmt.Println("server is starting")
	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", r))

}
