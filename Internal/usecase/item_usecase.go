package usecase

import (
	"errors"
	"fmt"
	"meli-items-w2/Internal/entity"
)

type ItemUsecase interface {
	AddItem(item entity.Item) (*entity.Item, error)
	UpdateItemById(item entity.Item, id int) (*entity.Item, error)
	GetItemById(id int) *entity.Item
	DeleteItemById(id int) *entity.Item
	GetAllItems(status string, limit int) []entity.Item
}

type itemUsecase struct {
	repo entity.ItemRepository
}

func NewItemUsecase(repo *interface{}) *itemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) AddItem(item entity.Item) (entity.Item, error) {
	items := u.repo.GetItems()
	for _, b := range items {
		if b.Code == item.Code {
			return entity.Item{}, errors.New("Item already exist")
		}
	}
	err := u.repo.AddItem(&item)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}

	return item, nil

}

func (u *itemUsecase) UpdateItemById(item entity.Item, id int) (entity.Item, error) {
	items := u.repo.GetItems()
	for _, b := range items {
		if b.Code == item.Code && b.Id == id {
			return entity.Item{}, errors.New("Item already exist")
		}
	}
	err := u.repo.AddItem(&item)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}
	return item, nil
}

func (u *itemUsecase) GetItemByID(id int) *entity.Item {
	return u.repo.GetItem(id)
}

func (u *itemUsecase) DeleteItemById(id int) *entity.Item {
	return u.repo.DeleteItem(id)
}

func (u *itemUsecase) GetAllItems() []entity.Item {
	return u.repo.GetItems()
}
