package usecase

import (
	"errors"
	"fmt"
	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/domain/port"
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

//definimos una estructura
type itemUsecase struct {
	repo port.ItemRepository
}

//funcion new
func NewItemUsecase(repo port.ItemRepository) port.ItemUsecase {
	return itemUsecase{
		repo: repo,
	}
}

/* --------GET --------*/
func (u itemUsecase) Index() []domain.Item {
	return u.repo.Index()
}

func (u itemUsecase) GetListaInicial() []domain.Item{
	return u.repo.GetListaInicial()
}

func (u itemUsecase) GetAllItems() []domain.Item{
	return u.repo.GetAllItems()
}

func(u itemUsecase) GetItemById(id int) *domain.Item {
	return u.repo.GetItemById(id)
}


/*------------POST-----------*/
func(u itemUsecase) AddItem(item domain.Item) (*domain.Item, error){
	items := u.repo.GetAllItems()
	for _, b := range items{
		if b.ID == item.ID {
			return nil, errors.New("item already exists")
		}
	}
	result := u.repo.AddItem(item)
	return result, nil
}

/*---------PUT----------------*/
func(u itemUsecase) UpdateItem(item domain.Item) (*domain.Item, error) {
	//validaciones y logica del negocio
	// validación de código único, no va codigo repetido
	//se niega
	almacenaje := u.repo.GetItemById(item.ID)
	if almacenaje == nil {
		return almacenaje, fmt.Errorf("el item no existe")
	}
	u.ActualizarUpdateAt(&item)
	err := u.repo.UpdateItemNuevo(item)
	if err != nil{
return &item,fmt.Errorf("error actualizando item")
	}
	return &item,nil
}

func (u itemUsecase) ActualizarUpdateAt(item *domain.Item){

}


func (u itemUsecase)UpdateItemNuevo(item domain.Item){

}