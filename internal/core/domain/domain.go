package domain

import "time"

type Item struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Stock       uint      `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/*
type ItemRepository interface {
	AddItem(item Item) *Item
	UpdateItem(item Item, id int) *Item
	GetItemById(id int) *Item
	DeleteItem(id int) *Item
	ListItem(status string) []Item

	GetNextId() int
	GetItemByCode(code string) *Item

	GetDB() []Item
}
*/
