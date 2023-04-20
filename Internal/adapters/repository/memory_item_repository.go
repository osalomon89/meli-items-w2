package repository

import (
	"errors"
	"meli-items-w2/Internal/entity"
	"time"
)

var newID uint = 0

type itemRepository struct {
	db []entity.Item
}

func NewItemRepository() entity.ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetItems() []entity.Item {
	return r.db
}

func (r *itemRepository) GetItem(id uint) (entity.Item, error) {

	for _, item := range r.db {
		if uint(item.Id) == id {
			return item, nil
		}
	}

	return entity.Item{}, errors.New("Item not found")
}

func (r *itemRepository) AddItem(item *entity.Item) error {

	createdAt := time.Now()
	newID = newID + 1

	item.Id = int(newID)
	item.CreatedAt = createdAt
	item.UpdatedAt = createdAt
	r.db = append(r.db, *item)

	return nil

	//r.db = append(r.db, item)
	//return &item
}
