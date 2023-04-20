package usecase //Aplicacion Bussiness Rules  (Capa de aplicacion)

//Solo conoce las entidades (dominio, eventos, servicios)
import (
	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
	//"errors"
)

// Interface que recibe los metodos
type ItemUseCase interface {
	GetItems() ([]dom.Item, error)
	//AddItems(item dom.Item) *dom.Item
	//GetItemsById(item dom.Item) *dom.Item
	//UpdateItems(item dom.Item) *dom.Item
	//DeleteItem(item dom.Item) *dom.Item //luego seria ideal retornar el error tambine (*dom.Item, error)

}

// Db
type itemUseCase struct {
	repo dom.ItemRepository
}

// Contructor que recibe ItemUseCase
func NewItemUseCase(repo dom.ItemRepository) ItemUseCase {
	return &itemUseCase{
		repo: repo,
	}
}

func (uc *itemUseCase) GetItems() ([]dom.Item, error) {
	return uc.repo.GetDB(), nil
}
