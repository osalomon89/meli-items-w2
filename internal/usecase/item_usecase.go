package usecase

import (
	"errors"
	"gigigarino/challengeMELI/internal/domain"
)

/*
en esta capa definimos una estructura y una interface con los metodos o casos de uso
getallbooks
getbookporid
addbook
y retornamos el newbook usecase

se definen todas las funciones
*/

//definimos una interface

type ItemUsecase interface {
	//metodos 
	GetAllItems() []domain.Item
	GetItemById(id int) *domain.Item
	//GetListaInicial() []domain.Item
	AddItem (item domain.Item)(*domain.Item, error)
}


//definimos una estructura
type itemUsecase struct {
	repo domain.ItemRepository
}

//funcion new
func NewItemUsecase(repo domain.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) GetAllItems() []domain.Item{
	return u.repo.GetAllItems()
}

func(u *itemUsecase) GetItemById(id int) *domain.Item {
	return u.repo.GetItemById(id)
}

func(u* itemUsecase) AddItem(item domain.Item) (*domain.Item, error){
	items := u.repo.GetAllItems()
	for _, b := range items{
		if b.ID == item.ID {
			return nil, errors.New("item already exists")
		}
	}
	result := u.repo.AddItem(item)
	return result, nil
}