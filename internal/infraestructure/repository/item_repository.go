package repository

import (
	"gigigarino/challengeMELI/internal/domain"
	
)

//definimos estructura
type itemRepository struct {
	articulos []domain.Item
}

//definimos nuevafunciom
func NewItemRepository() domain.ItemRepository {
	return &itemRepository{}
}

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

func (r * itemRepository) AddItem (item domain.Item) *domain.Item{
	r.articulos = append(r.articulos, item)
	return &item
}