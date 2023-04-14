package usecase

import (
	"errors"
	"fmt"

	"github.com/osalomon89/meli-items-w2/internal/domain"
)

type ItemUsecase interface {
	AddItem(item domain.Item) (*domain.Item, error)
	UpdateItemById(item domain.Item, id int) (*domain.Item, error)
	GetItemById(id int) *domain.Item
	DeleteItemById(id int) *domain.Item
	GetAllItems(status string, limit int) []domain.Item
}

type itemUsecase struct {
	repo domain.ItemRepository
}

func NewItemUsecase(repo domain.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (uc *itemUsecase) AddItem(item domain.Item) (*domain.Item, error) {
	if err := uc.validateCode(item.Code); err != nil {
		return nil, err
	}
	return uc.repo.AddItem(item), nil
}

func (uc *itemUsecase) UpdateItemById(item domain.Item, id int) (*domain.Item, error) {
	if err := uc.validateCodeUpdate(item); err != nil {
		return nil, err
	}
	return uc.repo.UpdateItem(item, id), nil
}

func (uc *itemUsecase) GetItemById(id int) *domain.Item {
	return uc.repo.GetItem(id)
}

func (uc *itemUsecase) DeleteItemById(id int) *domain.Item {
	return uc.repo.DeleteItem(id)
}

func (uc *itemUsecase) GetAllItems(status string, limit int) []domain.Item {
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

func (uc *itemUsecase) validateCodeUpdate(item domain.Item) error {

	for _, v := range uc.repo.GetDB() {
		if v.Code == item.Code && v.Id != item.Id {
			return errors.New(fmt.Sprintf("The code '%s' already exists", item.Code))
		}
	}
	return nil
}
