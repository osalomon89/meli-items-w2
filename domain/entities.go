package domain

import(
	"time"
)

type Item struct {
	ID     int    `json:"id"`
	Code string `json:"code"`
	Title  string `json:"title"`
	Description  string    `json:"description"`
	Price     int    `json:"price"`
	Stock int `json:"stock"`
	Status  string `json:"status"`
	Photos string `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}