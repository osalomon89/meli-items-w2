package repository

import (
	"errors"
	"math/rand"
	"time"

	dom "meli-items-w2/domain"
)

var repo itemRepository

type itemRepository struct {
	itemsDB []dom.Item
	countId int
}

// AddItem Añadir item
func (iRepo itemRepository) AddItem(item dom.Item) *dom.Item {
	//TODO verificar datos de entrada y devolver nil, posiblemente encerrar todo en un if

	// Id Incremental
	iRepo.countId++
	item.Id = iRepo.countId

	// CreatedAt toma el momento de creación, UpdatedAt toma hora 0 por default
	item.CreatedAt = time.Now()

	// Asignamos status de acuerdo al stock
	item.Status = statusCheck(item.Stock)
	// Verificamos el código único
	item.Code = codeCheck(item.Code)

	// Guardo
	iRepo.itemsDB = append(iRepo.itemsDB, item)

	return &item
}

// GetItemById Obtener item por ID
func (iRepo *itemRepository) GetItemById(id int) *dom.Item {
	for _, item := range iRepo.itemsDB {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

// UpdateItem modificar item
func (iRepo itemRepository) UpdateItem(item dom.Item, id int) *dom.Item {
	itemFound := iRepo.GetItemById(id)

	if itemFound == nil {
		return nil
	}

	itemFound.Code = codeCheck(item.Code)
	itemFound.Title = item.Title
	itemFound.Description = item.Description
	itemFound.Price = item.Price
	itemFound.Stock = item.Stock
	itemFound.Status = statusCheck(item.Stock)
	itemFound.UpdatedAt = time.Now()

	validateItem(item)

	for i, v := range iRepo.itemsDB {
		if v.Id == itemFound.Id {
			iRepo.itemsDB[i] = *itemFound
			return itemFound
		}
	}

	return nil
}

func (iRepo itemRepository) DeleteItem(id int) *dom.Item {
	//TODO implement me
	panic("implement me")
}

func (iRepo itemRepository) ListItem(status string, limit int) []dom.Item {
	//TODO implement me
	panic("implement me")
}

// Obtener base
func (iRepo itemRepository) GetDB() []dom.Item {
	return iRepo.itemsDB
}

func NewItemRepository() dom.ItemRepository {
	return &itemRepository{}
}

// ******** FUNCIONES AUXILIARES ********

func statusCheck(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}

func codeCheck(code string) string {
	// Sólo en caso tal de que nos quedemos atrapados en un bucle
	attempts := 0

	// Bucle para que checkee cuantas veces sea necesario
	for {
		// Si codeFound está en en la bd, rompe y sale a generar un código
		// el bucle for vuelve a comenzar, y la variable codeFound inicia en false
		// si esta vez es diferente, sale por el if == false
		codeFound := false
		for _, item := range repo.itemsDB {
			if item.Code == code {
				codeFound = true
				break
			}
		}
		if codeFound == false {
			return code
		}
		attempts++
		if attempts > 10 {
			break
		}
		code = generateCode()
	}
	return ""
}

func generateCode() string {
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

func validateItem(item dom.Item) error {

	if item.Title == "" || item.Description == "" {
		return errors.New("tittle or description are required")
	}

	if item.Price < 0 || item.Stock < 0 {
		return errors.New("price or stock should be greater than 0")
	}

	if item.Code == "" || len(item.Code) != 11 {
		return errors.New("code is not valid")
	}

	return nil
}
