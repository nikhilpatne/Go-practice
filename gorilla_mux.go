package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	ISBN  string `json:"isbn"`
	// Description string `json:"description"`
	// Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {

	r := mux.NewRouter()

	// Initialize struct

	books = append(books, Book{ID: "1", ISBN: "123455", Title: "Wings of fire", Author: &Author{Firstname: "Nikhil", Lastname: "Patne"}})

	books = append(books, Book{ID: "2", ISBN: "112314", Title: "nihil patne", Author: &Author{Firstname: "Mahesh", Lastname: "Pol"}})

	r.HandleFunc("/books", getAllbooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9092", r))
}

func getAllbooks(w http.ResponseWriter, r *http.Request) {
	// set header

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, items := range books {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}

	json.NewEncoder(w).Encode(Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
