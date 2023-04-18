package entity

import (
	"fmt"
	"time"
)

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

type ItemRepository interface {
	AddItem(item *Item) error
	UpdateItem(item *Item, id int) (Item, error)
	GetItem(id int) (Item, error)
	DeleteItem(id int) (Item, error)
	GetItems(status string, limit int) ([]Item, error)
	ValidateCode(code string) (bool, error)
	ValidateCodeUpdate(item Item, id int) (bool, error)
	ValidateStatus(int) string
}

type ItemAlreadyExist struct {
	Message string
}

func (e ItemAlreadyExist) Error() string {
	return fmt.Sprintf("error: '%s'", e.Message)
}

type ItemNotFound struct {
	Message string
}

func (i ItemNotFound) Error() string {
	return i.Message
}
