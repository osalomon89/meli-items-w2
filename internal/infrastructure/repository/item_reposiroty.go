package repository

import (
	"errors"

	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
)

type itemRepository struct {
	db []dom.Item
}

func NewItemRepository() dom.ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetDB() []dom.Item {
	return r.db
}

/* func (r *itemRepository) GetItem(id int) *dom.Item {
	for _, v := range r.db {
		if v.ID == id {
			return &v
		}
	}
	return nil
}  */

// funciones que ayudan a las principales :v

func (r *itemRepository) ChangeItemStatus(item *dom.Item) error {
	if item == nil {
		return errors.New("item is nil")
	}
	if item.Stock == 0 {
		item.Status = "INACTIVE"
		return nil
	}

	item.Status = "ACTIVE"
	return nil
}

func (r *itemRepository) FindItemById(id int) *dom.Item {
	for i := range r.db {
		if r.db[i].ID == id {
			return &r.db[i]
		}
	}
	return nil
}

func (r *itemRepository) GenerateID(items []dom.Item) int {
	maxId := 0
	for i := 0; i < len(items); i++ {
		if items[i].ID > maxId {
			maxId = items[i].ID
		}
	}
	return maxId + 1
}

func (r *itemRepository) RequeriedFields(item *dom.Item) error {
	if item == nil {
		return errors.New("item is nil")
	}
	if item.Code == "" {
		return errors.New("code is required")
	}
	if item.Title == "" {
		return errors.New("title is required")
	}
	if item.Description == "" {
		return errors.New("description is required")
	}
	if item.Price == 0 || item.Price < 0 {
		return errors.New("price is required and need be greater that 0")

	}
	if item.Stock < 0 {
		return errors.New("stock need be greater that 0")
	}
	return nil
}

func (r *itemRepository) UpdateFields(item *dom.Item, updateItem dom.Item) {

	if updateItem.Code != "" {
		item.Code = updateItem.Code
	}
	if updateItem.Title != "" {
		item.Title = updateItem.Title
	}
	if updateItem.Description != "" {
		item.Description = updateItem.Description
	}
	if updateItem.Price != 0 {
		item.Price = updateItem.Price
	}
	if updateItem.Stock != 0 {
		item.Stock = updateItem.Stock
	}
}

func (r *itemRepository) VerifyCode(code string) bool {
	for i := range r.db {
		if r.db[i].Code == code {
			return false
		}
	}
	return true
}
