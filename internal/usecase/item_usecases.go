package usecase

import (
	"errors"
	"fmt"

	"github.com/osalomon89/meli-items-w2/internal/entity"
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

func NewItemUsecase(repo entity.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (uc *itemUsecase) AddItem(item entity.Item) (*entity.Item, error) {
	if err := uc.validateCode(item.Code); err != nil {
		return nil, err
	}
	return uc.repo.AddItem(item), nil
}

func (uc *itemUsecase) UpdateItemById(item entity.Item, id int) (*entity.Item, error) {
	if err := uc.validateCodeUpdate(item, id); err != nil {
		return nil, err
	}
	return uc.repo.UpdateItem(item, id), nil
}

func (uc *itemUsecase) GetItemById(id int) *entity.Item {
	return uc.repo.GetItem(id)
}

func (uc *itemUsecase) DeleteItemById(id int) *entity.Item {
	return uc.repo.DeleteItem(id)
}

func (uc *itemUsecase) GetAllItems(status string, limit int) []entity.Item {
	return uc.repo.GetItems(status, limit)
}

// *****************Funciones auxiliares*****************
func (uc *itemUsecase) validateCode(code string) error {

	for _, v := range uc.repo.GetDB() {
		if v.Code == code {
			return errors.New(fmt.Sprintf("The code '%s' already exists", code))
		}
	}
	return nil
}

func (uc *itemUsecase) validateCodeUpdate(item entity.Item, id int) error {

	for _, v := range uc.repo.GetDB() {
		if v.Code == item.Code && v.Id != id {
			return errors.New(fmt.Sprintf("The code '%s' already exists", item.Code))
		}
	}
	return nil
}
