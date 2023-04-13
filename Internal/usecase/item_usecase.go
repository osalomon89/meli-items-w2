package usecase

import (
	"errors"
	"meli-items-w2/Internal/domain"
)

type ItemUsecase interface {
	GetAllBooks() []domain.Item
	GetBookByID(id int) *domain.Item
	AddBook(book domain.Item) (*domain.Item, error)
}

type itemUsecase struct {
	repo domain.BookRepository
}

func NewItemUsecase() ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) GetAllItems() []domain.Item {
	return u.repo.GetItmes()
}

func (u *itemUsecase) GetItemByID(id int) *domain.Item {
	return u.repo.GetItem(id)
}

func (u *itemUsecase) AddItem(book domain.Item) (*domain.Item, error) {
	books := u.repo.GetItems()
	for _, i := range books {
		if i.ID == Item.ID {
			return nil, errors.New("book already exist")
		}
	}

	result := u.repo.AddItem(book)

	return result, nil
}
