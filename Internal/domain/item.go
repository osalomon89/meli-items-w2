package domain

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
	Author      string  `json:"author"`
}
