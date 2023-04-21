package port

import (
	"gigigarino/challengeMELI/internal/domain"
)

//go:generate mockgen -source=./ports.go -destination=../../mocks/item_usecase_mock.go -package=mocks
//mockgen -source=item-repository.go -destination=./mocks/item_repository_mock.go -package=mocks

type ItemUsecase interface {
	//metodos 
	
	Index() []domain.Item
	GetAllItems() []domain.Item
	GetItemById(int) *domain.Item
	GetListaInicial() []domain.Item
	AddItem(domain.Item)(*domain.Item, error)
	UpdateItem(domain.Item)(*domain.Item, error)
	ActualizarUpdateAt(*domain.Item)
	UpdateItemNuevo(domain.Item)
	DeleteItem(int) (bool, error)
}


type ItemRepository interface {
	/*---------GETS---------*/
	Index() []domain.Item
	GetListaInicial() []domain.Item
	GetAllItems() []domain.Item
	GetItemById(int)*domain.Item

	/*---------POSTS---------*/
	AddItem(domain.Item)*domain.Item

	/*---------PUTS---------*/
	//UpdateItem(id int)*domain.Item
	UpdateItemNuevo(domain.Item) error

	//--------DELETE----------*/
	DeleteItem(int) (bool)
	
}

