package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"auhor"`
}

// Author model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lirstname"`
}

var books []Book

// ? get all
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// ? get by id
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	//Loop trough books to find ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// ? POST
func createBook(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}

// ^ PUT (UPDATE)
func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(res).Encode(book)
			return
		}
	}
}

// ! DELETE
func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "982345", Title: "Solaris", Author: &Author{Firstname: "Stanislaw", Lastname: "Lem"}})
	books = append(books, Book{ID: "2", Isbn: "233456", Title: "Dune", Author: &Author{Firstname: "Frank", Lastname: "Herbert"}})
	books = append(books, Book{ID: "3", Isbn: "569032", Title: "Aboba", Author: &Author{Firstname: "Jmishenko", Lastname: "Valera"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	fmt.Println("http://localhost:8000/api/books")
	log.Fatal(http.ListenAndServe(":8000", router))
}
