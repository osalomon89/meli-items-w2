package repository

import (
	"math/rand"
	"time"

	dom "meli-items-w2/domain"
)

var countId = 0

var repo itemRepository

type itemRepository struct {
	itemsDB []dom.Item
}

// Añadir item
func (iRepo itemRepository) AddItem(item dom.Item) *dom.Item {
	// Id Incremental
	countId++
	item.Id = countId

	// CreatedAt toma el momento de creación, UpdatedAt toma hora 0
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Time{}

	// Asignamos status de acuerdo al stock
	item.Status = statusCheck(item.Stock)
	item.Code = codeCheck(item.Code)

	// Guardo
	iRepo.itemsDB = append(iRepo.itemsDB, item)

	return &item
}

func (iRepo itemRepository) UpdateItem(item dom.Item, id int) *dom.Item {
	//TODO implement me
	panic("implement me")
}

func (iRepo itemRepository) GetItem(id int) *dom.Item {
	//TODO implement me
	panic("implement me")
}

func (iRepo itemRepository) DeleteItem(id int) *dom.Item {
	//TODO implement me
	panic("implement me")
}

func (iRepo itemRepository) GetItems(status string, limit int) []dom.Item {
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
		code = generateCode()
	}
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
