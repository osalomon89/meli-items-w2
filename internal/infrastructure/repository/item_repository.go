package repository

import (
	"errors"
	"fmt"

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
func (r *itemRepository) SaveItem(item *dom.Item) error {
	r.db = append(r.db, *item)
	return nil
}

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

func (r *itemRepository) GenerateID() int {
	var nextId int
	for _, val := range r.db {
		if nextId < val.ID {
			nextId = val.ID
		}
	}
	nextId++
	return nextId
}

func (r *itemRepository) RequiredFields(item *dom.Item) error {
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

/*
	 func (r *itemRepository) UpdateFields(item *dom.Item) error {
		for i := range r.db {
			if r.db[i].ID == item.ID {
				// Actualizar el item en el slice
				r.db[i] = *item
				return nil
			}
		}
		return fmt.Errorf("item with id %d does not exist", item.ID)
	}
*/
func (r *itemRepository) UpdateFields(updateItem *dom.Item) error {
	for i := range r.db {
		if r.db[i].ID == updateItem.ID {
			if updateItem.Code != "" {
				r.db[i].Code = updateItem.Code
			}
			if updateItem.Title != "" {
				r.db[i].Title = updateItem.Title
			}
			if updateItem.Description != "" {
				r.db[i].Description = updateItem.Description
			}
			if updateItem.Price != 0 {
				r.db[i].Price = updateItem.Price
			}
			if updateItem.Stock != 0 {
				r.db[i].Stock = updateItem.Stock
			}
			return nil
		}
	}
	return fmt.Errorf("item with id %d does not exist", updateItem.ID)
}

func (r *itemRepository) VerifyCode(code string) bool {
	for i := range r.db {
		if r.db[i].Code == code {
			return true
		}
	}
	return false
}

func (r *itemRepository) DeleteRegister(id int, item *dom.Item) error {
	if item == nil {
		return fmt.Errorf("item not exist (Error en respository - deleteRegister)")
	}
	for i, item := range r.db {
		if item.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			break
		}
	}
	return nil
}
