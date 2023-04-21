package usecase //Aplicacion Bussiness Rules  (Capa de aplicacion)

//Solo conoce las entidades (dominio, eventos, servicios)
import (
	"errors"
	"fmt"
	"time"

	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
	//"errors"
)

// Interface que recibe los metodos
type ItemUseCase interface {
	GetItems() ([]dom.Item, error)
	AddItem(item *dom.Item) error
	GetItemById(id int) (*dom.Item, error)
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

func (uc *itemUseCase) AddItem(item *dom.Item) error {

	// Verificar que el codigo sea unico
	if uc.repo.VerifyCode(item.Code) {
		return fmt.Errorf("item with code %s already exists", item.Code)
	}
	// Campos requeridos
	if err := uc.repo.RequiredFields(item); err != nil {
		return err
	}
	// Cambiar status segun el stock
	if err := uc.repo.ChangeItemStatus(item); err != nil {
		return fmt.Errorf("error changing item status: %w", err)
	}

	//generar id nuevo
	newId := uc.repo.GenerateID()
	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt
	item.ID = newId
	if err := uc.repo.SaveItem(item); err != nil {
		return err
	}

	return nil
}

func (c *itemUseCase) GetItemById(id int) (*dom.Item, error) {
	item := c.repo.FindItemById(id)
	if item == nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}
