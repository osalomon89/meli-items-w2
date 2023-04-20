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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	//CreatAt     string  `json:"creat_at"`
	//UpdateAt    string  `json:"update_at"`

}
type ItemRepository interface {
	AddItem(item *Item) error
	//UpdateItem(item Item, id int) *Item
	GetItem(id uint) (Item, error)
	//DeleteItem(id int) *Item
	GetItems() []Item
}
