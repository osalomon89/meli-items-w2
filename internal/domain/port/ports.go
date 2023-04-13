package port

import (
	"gigigarino/challengeMELI/internal/domain"
)


type ItemUsecase interface {
	//metodos 
	
	Index() []domain.Item
	GetAllItems() []domain.Item
	GetItemById(id int) *domain.Item
	GetListaInicial() []domain.Item
	AddItem(item domain.Item)(*domain.Item, error)
	UpdateItem(item domain.Item)(*domain.Item, error)
	ActualizarUpdateAt(item *domain.Item)
	UpdateItemNuevo(item domain.Item)
}


type ItemRepository interface {
	/*---------GETS---------*/
	Index() []domain.Item
	GetListaInicial() []domain.Item
	GetAllItems() []domain.Item
	GetItemById(id int)*domain.Item

	/*---------POSTS---------*/
	AddItem(item domain.Item)*domain.Item

	/*---------PUTS---------*/
	//UpdateItem(id int)*domain.Item
	UpdateItemNuevo(item domain.Item) error

	//auxiliares
	//ActualizarUpdateAt(item domain.Item)
	
}

