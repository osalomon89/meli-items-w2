package item_domain

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

type ItemRepository interface {
	AddItem(item Item) error
	UpdateItem(item Item) error
	GetById(id int) (*Item, error)
	DeleteById(id int) error
	GetByStatusAndLimit(status string, limit int) []Item
}
