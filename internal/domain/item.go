package domain //Tambien puedo llamar a la carpeta "ENTITIES" (Enterprise Bussiness Rules)

import (
	"time"
)

// Articulos   (las claves del json se obtienen en minusculas como "buena practicas")

// Esto es como una clase ()
type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type ItemRepository interface {
	//DeleteItem(item Item) bool
	GetDB() []Item
	GenerateID() int
	VerifyCode(code string) bool
	FindItemById(id int) *Item
	ChangeItemStatus(item *Item) error
	RequiredFields(item *Item) error
	UpdateFields(item *Item, updateItem Item)
	SaveItem(item *Item) error
}
