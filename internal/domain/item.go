package domain

import (
	"time"
)

type Item struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemRepository interface {
	GetItem(id int) *Item
	DeleteItem(id int) bool
	GetDB() []Item
	CodeRepetido(item Item) bool
	ObtenerSiguienteID() int
	SaveItem(item Item)
	ModifyItem(id int,item Item)
}

/*
type BookRepository interface {
	GetBooks() []Book
	GetBook(id int) *Book
	AddBook(book Book) *Book
}
*/