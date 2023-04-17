package usecase

import (
	dom "meli-items-w2/internal/domain"
)

type ItemUsecase interface {
	AddItem(item dom.Item) *dom.Item
	GetItemById(id int) *dom.Item
	UpdateItem(item dom.Item, id int) *dom.Item
	DeleteItem(id int) *dom.Item
	ListItem(status string) []dom.Item
}

type itemUsecase struct {
	repo dom.ItemRepository
}

func NewItemUsecase(repo dom.ItemRepository) *itemUsecase {
	return &itemUsecase{repo}
}

func (uc *itemUsecase) AddItem(item dom.Item) *dom.Item {
	return uc.repo.AddItem(item)
}

func (uc *itemUsecase) UpdateItem(item dom.Item, id int) *dom.Item {
	return uc.repo.UpdateItem(item, id)
}

func (uc *itemUsecase) GetItemById(id int) *dom.Item {
	return uc.repo.GetItemById(id)
}

func (uc *itemUsecase) DeleteItem(id int) *dom.Item {
	return uc.repo.DeleteItem(id)
}

func (uc *itemUsecase) ListItem(status string) []dom.Item {
	return uc.repo.ListItem(status)
}
