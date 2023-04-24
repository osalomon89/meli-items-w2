//es la entidad y estructura

package domain

import "time"

type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

type ItemNotFound struct {
	Message string
}

//si quiero hacer un maoa
/*

type mapRepo map[int]Item
*/
