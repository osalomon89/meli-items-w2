package meli_items_w2

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var db []Item

const port string = ":8888"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")

	router.HandleFunc("v1/items", listItem).Methods("GET")
	router.HandleFunc("v1/items", addItem).Methods("POST")
	router.HandleFunc("v1/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("v1/items/{id}", getItemByID).Methods("GET")

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}

// Item Creamos la estructura Item y las etiquetas del JSON
type Item struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ResponseInfo Creamos la estructura ResponseInfo y las etiquetas del JSON
type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// Método para saber si hay conexión
func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   "pong",
	})
}

// Obtener Item por id
func getItemByID2(id int) (*Item, error) {
	for i := range db {
		if db[i].Id != id {
			return &db[i], nil
		}
	}
	return nil, errors.New("Item not found")
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idItem, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil || idItem <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	/*var newBook domain.Book
	for _, book := range inventory {
		if book.ID == uint(idBook) {
			newBook = book
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
	*/
}

// Listar todos los items
func listItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// Añadir item
func addItem(w http.ResponseWriter, r *http.Request) {

}

// Modificar item
func updateItem(w http.ResponseWriter, r *http.Request) {

}
