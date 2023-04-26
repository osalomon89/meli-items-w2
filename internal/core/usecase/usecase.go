package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	dom "meli-items-w2/internal/core/domain"
	port "meli-items-w2/internal/core/usecase/port"
)

type ItemUsecase interface {
	AddItem(item dom.Item) *dom.Item
	GetItemById(id int) *dom.Item
	UpdateItem(item dom.Item, id int) *dom.Item
	DeleteItem(id int) *dom.Item
	ListItem(status string) []dom.Item
}

type itemUsecase struct {
	repo port.ItemRepository
}

func NewItemUsecase(repo port.ItemRepository) *itemUsecase {
	return &itemUsecase{repo}
}

func (uc *itemUsecase) AddItem(item dom.Item) *dom.Item {
	// Id Incremental
	item.Id = uc.repo.GetNextId()

	// CreatedAt toma el momento de creación, UpdatedAt toma hora 0 por default
	item.CreatedAt = time.Now()

	// Asignamos status de acuerdo al stock
	item.Status = statusCheck(item.Stock)

	// Verificamos el código único
	item.Code = uc.codeCheck(item.Code)

	if uc.validateItem(item) == nil {
		uc.repo.AddItem(item)
		return &item
	}

	return nil

}

func (uc *itemUsecase) UpdateItem(item dom.Item, id int) *dom.Item {
	itemFound := uc.repo.GetItemById(id)

	if itemFound == nil {
		return nil
	}

	item.Code = uc.codeCheck(item.Code)
	itemFound.Title = item.Title
	itemFound.Description = item.Description
	itemFound.Price = item.Price
	itemFound.Stock = item.Stock
	itemFound.Status = statusCheck(item.Stock)
	itemFound.UpdatedAt = time.Now()

	if uc.validateItem(item) == nil {
		uc.repo.AddItem(item)
		return &item
	}

	updatedItem := uc.repo.UpdateItem(item, id)
	if updatedItem != nil {
		return updatedItem
	}
	return nil
}

func (uc *itemUsecase) GetItemById(id int) *dom.Item {
	return uc.repo.GetItemById(id)
}

func (uc *itemUsecase) DeleteItem(id int) *dom.Item {
	itemFound := uc.repo.GetItemById(id)

	if itemFound == nil {
		return nil
	}

	return uc.repo.DeleteItem(id)
}

func (uc *itemUsecase) ListItem(status string) []dom.Item {
	return uc.repo.ListItem(status)
}

// ---------> FUNCIONES VERIFICACIÓN <---------

func statusCheck(stock uint) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}

func (uc *itemUsecase) codeCheck(code string) string {
	// Sólo en caso tal de que nos quedemos atrapados en un bucle
	attempts := 0

	// Bucle para que checkee cuantas veces sea necesario
	for {
		// Si codeFound está en en la bd, rompe y sale a generar un código
		// el bucle for vuelve a comenzar, y la variable codeFound inicia en false
		// si esta vez es diferente, sale por el if == false

		codeFound := uc.repo.GetItemByCode(code)

		/*for _, item := range repo.itemsDB {
			if item.Code == code {
				codeFound = true
				break
			}
		}*/

		if codeFound == nil {
			return code
		}
		attempts++
		if attempts > 10 {
			break
		}
		code = uc.generateCode()
	}
	return ""
}

func (uc *itemUsecase) generateCode() string {
	// Crear un slice con los caracteres que quieres incluir
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generar un número aleatorio de 11 dígitos
	code := make([]byte, 11)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	// Retornar el código generado
	return string(code)
}

func (uc *itemUsecase) validateItem(item dom.Item) error {

	if item.Title == "" || item.Description == "" {
		fmt.Sprintf("tittle or description are required")
		return errors.New("tittle or description are required")
	}

	if item.Price < 0 || item.Stock < 0 {
		fmt.Sprintf("price or stock should be greater than 0")
		return errors.New("price or stock should be greater than 0")
	}

	if item.Code == "" || len(item.Code) != 11 {
		fmt.Sprintf("code is not valid")
		return errors.New("code is not valid")
	}

	return nil
}
