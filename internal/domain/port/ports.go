package port

import (
	"gigigarino/challengeMELI/internal/domain"
)


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

