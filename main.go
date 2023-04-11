package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const port = ":9000"

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

var db []Book

var errDB = errors.New("error in database")

func main() {
	b1 := Book{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	}

	b2 := Book{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	b3 := Book{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	}

	db = append(db, b1, b2, b3)

	r := mux.NewRouter()

	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

type ResourceNotFoundError struct {
	Message string
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		encodeJSON(w, ResponseInfo{
			Error: true,
			Data:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	key, err := findBook(id)
	if err != nil {
		notFoundError := new(ResourceNotFoundError)
		if ok := errors.As(err, notFoundError); ok {
			encodeJSON(w, ResponseInfo{
				Error: true,
				Data:  notFoundError.Error(),
			}, http.StatusNotFound)
			return
		}

		if ok := errors.Is(err, errDB); ok {
			encodeJSON(w, ResponseInfo{
				Error: true,
				Data:  errDB.Error(),
			}, http.StatusServiceUnavailable)
			return
		}

		encodeJSON(w, ResponseInfo{
			Error: true,
			Data:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	db[key] = book

	encodeJSON(w, db, http.StatusOK)
}

func findBook(id int) (int, error) {
	if id == 0 {
		return 0, errDB
	}

	if id == 99 {
		return 0, errors.New("internal server error")
	}

	for i, v := range db {
		if v.ID == id {
			return i, nil
		}
	}

	return 0, ResourceNotFoundError{
		Message: fmt.Sprintf("book not found. ID: %d", id),
	}
}

func encodeJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
