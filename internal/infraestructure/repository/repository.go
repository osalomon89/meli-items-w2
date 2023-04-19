package repository

import (
	dom "meli-items-w2/internal/domain"
)

type itemRepository struct {
	itemsDB []dom.Item
	countId int
}

// NewItemRepository Constructor
func NewItemRepository() dom.ItemRepository {
	return &itemRepository{}
}

func (iRepo *itemRepository) AddItem(item dom.Item) *dom.Item {
	iRepo.itemsDB = append(iRepo.itemsDB, item)
	return &item
}

func (iRepo *itemRepository) GetItemById(id int) *dom.Item {
	for _, item := range iRepo.itemsDB {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func (iRepo *itemRepository) UpdateItem(item dom.Item, id int) *dom.Item {
	for i, v := range iRepo.itemsDB {
		if v.Id == item.Id {
			iRepo.itemsDB[i] = item
			return &item
		}
	}
	return nil
}

func (iRepo *itemRepository) DeleteItem(id int) *dom.Item {

	for i, v := range iRepo.itemsDB {
		if v.Id == id {
			iRepo.itemsDB = append(iRepo.itemsDB[:i], iRepo.itemsDB[i+1:]...)
			return nil
		}
	}

	return nil

}

// TODO // limit: Es el tamaño solicitado de resultados en la página.
// TODO // Es un parámetro opcional, su valor default es 10, y su valor máximo es 20.

// ListItem listar bd aún no funcionan todos los filtros
func (iRepo *itemRepository) ListItem(status string) []dom.Item {
	// TODO filtrado por fecha y limite, agregar param en la interfaz y en la func ", limit int"

	var dbFiltered []dom.Item

	for _, item := range iRepo.itemsDB {
		if item.Status == status {
			dbFiltered = append(dbFiltered, item)
			return dbFiltered
		} else if status == "" {
			return iRepo.itemsDB
		}

	}

	return nil

}

func (iRepo *itemRepository) GetNextId() int {
	iRepo.countId++

	return iRepo.countId
}

func (iRepo *itemRepository) GetItemByCode(code string) *dom.Item {
	for _, item := range iRepo.itemsDB {
		if item.Code == code {
			return &item
		}
	}
	return nil
}

func (iRepo *itemRepository) GetDB() []dom.Item {
	return iRepo.itemsDB
}
