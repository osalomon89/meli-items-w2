package entity

import "time"

type Item struct {
	Id          int     `json:"id"`
	Code        string  `json:"code" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Descripcion string  `json:"descripcion" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Status      string  `json:"status"`
	CreatAt     string  `json:"creat_at"`
	UpdateAt    string  `json:"update_at"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type ItemRepository interface {
	AddItem(item Item) *Item
	UpdateItem(item Item, id int) *Item
	GetItem(id int) *Item
	DeleteItem(id int) *Item
	GetItems(status string, limit int) []Item
	GetDB() []Item
}
