package repository

import (
	dom "github.com/osalomon89/meli-items-w2/internal/domain"
)

type itemRepository struct {
	db []dom.Item
}

func NewItemRepository() dom.ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetDB() []dom.Item {
	return r.db
}

func (r *itemRepository) GetItem(id int) *dom.Item {
	for _, v := range r.db {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func (r *itemRepository) DeleteItem(id int) bool {
	var db_copy []dom.Item
	var encontrado bool

	for _,value := range r.db {
		if value.ID != id {
			db_copy = append(db_copy, value)
		} else {
			encontrado = true
		}
	}

	if encontrado {
		r.db = db_copy
	}

	return encontrado
}

func (r *itemRepository) CodeRepetido(id int, item dom.Item) bool {
	var repetido bool

	for _, val := range r.db {
		if id == 0 {
			if val.Code == item.Code {
				repetido = true
			}
		} else {
			if val.Code == item.Code && id != val.ID {
				repetido = true
			}
		}
	}
	return repetido
}

func (r *itemRepository) ObtenerSiguienteID() int {
	var idSiguiente int
	for _, val := range r.db {
		if idSiguiente < val.ID {
			idSiguiente = val.ID
		}
	}
	idSiguiente++
	return idSiguiente
}

func (r *itemRepository) SaveItem(item dom.Item){
	r.db = append(r.db, item)
}

func (r *itemRepository) ModifyItem(id int,item dom.Item) {
	for pos, val := range r.db {
		if val.ID == id {
			r.db[pos] = item
			return
		}
	}
}

func (r *itemRepository) GetItemsByStatus(status string) []dom.Item {
	var db_copy []dom.Item
	for _,value := range r.db {
		if value.Status == status {
			db_copy = append(db_copy, value)
		}
	}
	return db_copy
}
