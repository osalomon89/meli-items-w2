package port

import dom "meli-items-w2/internal/core/domain"

type ItemRepository interface {
	AddItem(item dom.Item) *dom.Item
	UpdateItem(item dom.Item, id int) *dom.Item
	GetItemById(id int) *dom.Item
	DeleteItem(id int) *dom.Item
	ListItem(status string) []dom.Item

	GetNextId() int
	GetItemByCode(code string) *dom.Item

	GetDB() []dom.Item
}
