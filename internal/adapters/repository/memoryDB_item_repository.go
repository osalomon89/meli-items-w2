package repository

import (
	"fmt"
	"sort"
	"time"

	"github.com/osalomon89/meli-items-w2/internal/entity"
)

type itemRepository struct {
	db []entity.Item
}

var countId int = 0

func NewItemRepository() entity.ItemRepository {
	return &itemRepository{}
}

func (repo *itemRepository) AddItem(item *entity.Item) error {
	countId++
	item.Id = countId

	item.Status = repo.ValidateStatus(item.Stock)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	repo.db = append(repo.db, *item)

	return nil
}

func (repo *itemRepository) UpdateItem(item *entity.Item, id int) (entity.Item, error) {
	for k, v := range repo.db {
		if v.Id == id {

			repo.db[k].Code = item.Code
			repo.db[k].Title = item.Title
			repo.db[k].Description = item.Description
			repo.db[k].Price = item.Price
			repo.db[k].Stock = item.Stock
			repo.db[k].Status = repo.ValidateStatus(item.Stock)
			repo.db[k].UpdatedAt = time.Now()

			return repo.db[k], nil
		}
	}

	return entity.Item{}, entity.ItemNotFound{
		Message: fmt.Sprintf("Item with id '%d' not found", id),
	}
}

func (repo *itemRepository) GetItem(id int) (entity.Item, error) {

	for _, v := range repo.db {
		if v.Id == id {
			return v, nil
		}
	}
	return entity.Item{}, entity.ItemNotFound{
		Message: fmt.Sprintf("Item with id '%d' not found", id),
	}

}

func (repo *itemRepository) DeleteItem(id int) error {
	for k, v := range repo.db {
		if v.Id == id {
			repo.db = append(repo.db[:k], repo.db[k+1:]...)
			return nil
		}
	}
	return entity.ItemNotFound{
		Message: fmt.Sprintf("Item with id '%d' not found", id),
	}
}

func (repo *itemRepository) GetItems(status string, limit int) ([]entity.Item, error) {

	if limit > 20 || limit > len(repo.db) || limit < 1 {
		limit = len(repo.db)
	}

	sort.SliceStable(repo.db, func(i, j int) bool {
		return repo.db[i].UpdatedAt.After(repo.db[j].UpdatedAt)
	})

	var itemsToshow []entity.Item

	if len(status) != 0 {
		for k, v := range repo.db {
			if v.Status == status {
				itemsToshow = append(itemsToshow, v)
			}
			if k == limit-1 {
				break
			}
		}
	} else {
		itemsToshow = append(repo.db[:limit])
	}

	if len(itemsToshow) == 0 {
		return itemsToshow, entity.ItemNotFound{
			Message: "None item matched the query",
		}
	}

	return itemsToshow, nil
}

func (repo *itemRepository) ValidateCode(code string) (bool, error) {
	for _, v := range repo.db {
		if v.Code == code {
			return true, nil
		}
	}
	return false, nil
}

func (repo *itemRepository) ValidateCodeUpdate(code string, id int) (bool, error) {

	for _, v := range repo.db {
		if v.Code == code && v.Id != id {
			return true, nil
		}
	}
	return false, nil
}

func (repo *itemRepository) ValidateStatus(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}
