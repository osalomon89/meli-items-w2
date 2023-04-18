package usecase

import (
	"fmt"

	"github.com/osalomon89/meli-items-w2/internal/entity"
)

type ItemUsecase interface {
	AddItem(item entity.Item) (entity.Item, error)
	UpdateItemById(item entity.Item, id int) (entity.Item, error)
	GetItemById(id int) (entity.Item, error)
	DeleteItemById(id int) (entity.Item, error)
	GetAllItems(status string, limit int) ([]entity.Item, error)
}

type itemUsecase struct {
	repo entity.ItemRepository
}

func NewItemUsecase(repo entity.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (uc *itemUsecase) AddItem(item entity.Item) (entity.Item, error) {
	isDuplicated, err := uc.repo.ValidateCode(item.Code)
	if err != nil {
		return item, fmt.Errorf("error in repository: %w", err)
	}
	if isDuplicated {
		return item, entity.ItemAlreadyExist{
			Message: fmt.Sprintf("Item already exists with code '%s", item.Code),
		}
	}
	if err = uc.repo.AddItem(&item); err != nil {
		return item, fmt.Errorf("error in repository: %w", err)
	}

	return item, nil
}

func (uc *itemUsecase) UpdateItemById(item entity.Item, id int) (entity.Item, error) {
	isDuplicated, err := uc.repo.ValidateCodeUpdate(item, id)
	if err != nil {
		return item, fmt.Errorf("error in repository: %w", err)
	}
	if isDuplicated {
		return item, entity.ItemAlreadyExist{
			Message: "Item already exists",
		}
	}

	result, err := uc.repo.UpdateItem(&item, id)
	if err != nil {
		return item, fmt.Errorf("error in repository: %w", err)
	}

	return result, nil
}

func (uc *itemUsecase) GetItemById(id int) (entity.Item, error) {
	item, err := uc.repo.GetItem(id)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}
	return item, nil
}

func (uc *itemUsecase) DeleteItemById(id int) (entity.Item, error) {
	item, err := uc.repo.DeleteItem(id)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}
	return item, nil
}

func (uc *itemUsecase) GetAllItems(status string, limit int) ([]entity.Item, error) {
	items, err := uc.repo.GetItems(status, limit)
	if err != nil {
		return []entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}
	return items, nil
}
