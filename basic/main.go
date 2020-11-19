package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/random"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get a single book 
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}