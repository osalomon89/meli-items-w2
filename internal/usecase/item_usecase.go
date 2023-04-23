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
	UpdateItem(item *dom.Item) (*dom.Item, error)
	DeleteItem(id int) error //luego seria ideal retornar el error tambine (*dom.Item, error)

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

func (uc *itemUseCase) GetItemById(id int) (*dom.Item, error) {
	item := uc.repo.FindItemById(id)
	if item == nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (uc *itemUseCase) UpdateItem(item *dom.Item) (*dom.Item, error) {
	if uc.repo.FindItemById(item.ID) == nil {
		return nil, fmt.Errorf("item with id %d does not exist", item.ID)
	}
	if err := uc.repo.ChangeItemStatus(item); err != nil {
		return nil, err
	}

	if err := uc.repo.UpdateFields(item); err != nil {
		return nil, err
	}

	item.UpdatedAt = time.Now()
	// Actualizar el item en el repositorio
	return item, nil
}

func (uc *itemUseCase) DeleteItem(id int) error {
	item := uc.repo.FindItemById(id)
	if item == nil {
		return fmt.Errorf("item with id %d does not exist", id)
	}
	if err := uc.repo.DeleteRegister(id, item); err != nil {
		return err
	}
	return nil
}
