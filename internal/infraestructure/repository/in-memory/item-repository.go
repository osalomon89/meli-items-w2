package repository

import (
	"fmt"
	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/domain/port"
	"time"
)

//definimos estructura como primer paso
//va en minuscula
type itemRepository struct {
	articulos []domain.Item
	
}

//definimos nuevafunciom
// en los aprametros de entrada hacemos la inyeccion de dependencias
func NewItemRepository() port.ItemRepository {
	return &itemRepository{}
}

// en el parametro de entrada se hace con puntero, se recomienda 
//son metodos ahora en vez de funciones 
func (r *itemRepository) Index() []domain.Item{
	return r.articulos
}

func (r *itemRepository) GetListaInicial() []domain.Item {
	return r.articulos
}

func (r *itemRepository) GetAllItems() []domain.Item {
	return r.articulos
}

func (r *itemRepository) GetItemById(id int) *domain.Item{
	for _, item := range r.articulos {
		if item.ID == id {
			return &item
		}
	}
	return nil
}



func (r * itemRepository) AddItem(item domain.Item) *domain.Item{
	r.articulos = append(r.articulos, item)
	return &item
}

//ver si esto esta bien
func (r *itemRepository) UpdateItem(id int) *domain.Item{
	// ver que se hace aca
	return nil
}







// funciones auxiliares 
//solo creacion podrian guardarse como save item
func (r *itemRepository) ActualizarCreateAt(item *domain.Item){
	item.CreatedAt = time.Now()
}

func (r *itemRepository) ActualizarUpdateAt(item *domain.Item){
	item.UpdateAt = time.Now()
}

func (r *itemRepository) ActualizarId(item *domain.Item){
	item.ID = r.ObtenerId()
}

//agregar
func (r *itemRepository)AppendItemToArticulos(item domain.Item){
	r.articulos = append(r.articulos, item)
}

//no puntero
//actualizar 
func (r *itemRepository)UpdateItemNuevo(item domain.Item)error{
	//le pasamos el puntero
	for i, v := range r.articulos {
		if v.ID == item.ID {
			r.articulos[i].Code = item.Code
			r.articulos[i].Title = item.Title
			r.articulos[i].Description = item.Description
			r.articulos[i].Price = item.Price
			r.articulos[i].Stock = item.Stock
			return nil
		}
	}
	return fmt.Errorf("el item no existe")
}


// codigo debe ser unico
// funcion repetido
func (r *itemRepository) CodigoRepetido(item *domain.Item) bool {
	var repetido bool
	//validar si es desigual a nill
	if item == nil{
		return true
	}
	for _, valor := range r.articulos {
		if valor.Code == item.Code {
			repetido = true
		}
	}
	return repetido
}

//id autoincremental, obtiene el proximo ID libre
func (r *itemRepository) ObtenerId() int {
	var idSiguiente int
	for _, itemAnterior := range r.articulos {
		if idSiguiente < itemAnterior.ID {
			idSiguiente = itemAnterior.ID
		}
	}
	idSiguiente += 1
	return idSiguiente
}

// validar status
func (r *itemRepository) ValidateStatus(item *domain.Item) {
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
}

func (r *itemRepository) InformacionCompleta(item *domain.Item) error{
	if item.Code == "" || item.Title == "" || item.Description == "" || item.Price == 0 {
		return fmt.Errorf("el campo no debe estar vacio")
	}
	return nil
}


//-------detele-----
func (r *itemRepository) DeleteItem(id int) bool {
	for i, v := range r.articulos {
		if v.ID == id {
			r.articulos = append(r.articulos[:i], r.articulos[i+1:]...)
			return true
		}
	}
	return false
}