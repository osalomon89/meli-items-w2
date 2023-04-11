package model

type Item struct {
	Id          int
	Code        string
	Title       string
	Description string
	Price       int
	Stock       int
	Status      string
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
