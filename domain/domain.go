package domain

import "time"

type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}
