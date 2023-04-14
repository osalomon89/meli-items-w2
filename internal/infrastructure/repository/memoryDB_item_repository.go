package repository

import (
	"sort"
	"time"

	"github.com/osalomon89/meli-items-w2/internal/domain"
)

type itemRepository struct {
	db []domain.Item
}

var countId int = 0

func NewItemRepository() domain.ItemRepository {
	return &itemRepository{}
}

func (repo *itemRepository) AddItem(item domain.Item) *domain.Item {
	countId++
	item.Id = countId

	item.Status = validateStatus(item.Stock)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	repo.db = append(repo.db, item)

	return &item
}

func (repo *itemRepository) UpdateItem(item domain.Item, id int) *domain.Item {
	for _, v := range repo.db {
		if v.Id == id {

			v.Code = item.Code
			v.Title = item.Title
			v.Description = item.Description
			v.Price = item.Price
			v.Stock = item.Stock
			v.Status = validateStatus(item.Stock)
			v.UpdatedAt = time.Now()

			return &item
		}
	}

	return nil
}

func (repo *itemRepository) GetItem(id int) *domain.Item {

	for _, v := range repo.db {
		if v.Id == id {
			return &v
		}
	}
	return nil

}

func (repo *itemRepository) DeleteItem(id int) *domain.Item {
	for k, v := range repo.db {
		if v.Id == id {
			repo.db = append(repo.db[:k], repo.db[k+1:]...)
			return &v
		}
	}
	return nil
}

func (repo *itemRepository) GetItems(status string, limit int) []domain.Item {

	if limit <= 0 {
		limit = 10
	} else if limit > 20 {
		limit = 20
	} else if limit > len(repo.db) {
		limit = len(repo.db)
	}

	sort.SliceStable(repo.db, func(i, j int) bool {
		return repo.db[i].UpdatedAt.After(repo.db[j].UpdatedAt)
	})

	var itemsToshow []domain.Item

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

	return itemsToshow
}

func (repo *itemRepository) GetDB() []domain.Item {
	return repo.db
}

// *****************Funcion auxiliares*****************
func validateStatus(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}
