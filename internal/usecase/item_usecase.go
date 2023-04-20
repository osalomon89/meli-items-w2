package usecase //Aplicacion Bussiness Rules  (Capa de aplicacion)

//Solo conoce las entidades (dominio, eventos, servicios)
import (
	"fmt"
	"time"

	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
	//"errors"
)

// Interface que recibe los metodos
type ItemUseCase interface {
	GetItems() ([]dom.Item, error)
	AddItems(item *dom.Item) (*dom.Item, error)
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

func (uc *itemUseCase) AddItems(item *dom.Item) (*dom.Item, error) {

	// Verificar que el codigo sea unico
	if !uc.repo.VerifyCode(item.Code) {
		return nil, fmt.Errorf("item with code %s already exists", item.Code)
	}
	// Campos requeridos
	if err := uc.repo.RequiredFields(item); err != nil {
		return nil, fmt.Errorf("required fields are missing: %w", err)
	}
	// Cambiar status segun el stock
	if err := uc.repo.ChangeItemStatus(item); err != nil {
		return nil, fmt.Errorf("error changing item status: %w", err)
	}
	db := uc.repo.GetDB()
	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt
	newId := uc.repo.GenerateID(db)
	item.ID = newId
	uc.repo.SaveItem(*item)

	return item, nil
}
