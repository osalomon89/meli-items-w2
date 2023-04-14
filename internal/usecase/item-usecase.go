package usecase

import (
	"fmt"
	"time"
	"sort"

	dom "github.com/osalomon89/meli-items-w2/internal/domain"
)

type ItemUsecase interface {
	GetItemByID(id int) *dom.Item
	DeleteItemByID(id int) bool
	GetAllItems() []dom.Item
	AddItemByItem(item *dom.Item) (*dom.Item,error)
	setStatus(item *dom.Item)
	UpdateItemByItem(id int,item *dom.Item) (*dom.Item,error)
	GetItemsByStatusAndLimit(status string, limit int) ([]dom.Item,error)
}

type itemUsecase struct {
	repo dom.ItemRepository
}

func NewItemUsecase(repo dom.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) GetAllItems() []dom.Item {
	return u.repo.GetDB()
}

func (u *itemUsecase) GetItemByID(id int) *dom.Item {
	return u.repo.GetItem(id)
}

func (u *itemUsecase) DeleteItemByID(id int) bool {
	return u.repo.DeleteItem(id)
}

func (u *itemUsecase) AddItemByItem(item *dom.Item) (*dom.Item,error) {
	if item == nil {
		return	nil,fmt.Errorf("invalid item")
	} 

	if u.repo.CodeRepetido(0,*item) {
		return	nil,fmt.Errorf("invalid code")
	}

	item.ID = u.repo.ObtenerSiguienteID()
	dt := time.Now()
	item.CreatedAt = dt
	item.UpdatedAt = dt
	u.setStatus(item)
	
	u.repo.SaveItem(*item)
	return item,nil
}

func (u *itemUsecase) UpdateItemByItem(id int, item *dom.Item) (*dom.Item,error){
	if item == nil {
		return	nil,fmt.Errorf("invalid item")
	} 
	if u.repo.CodeRepetido(id,*item) {
		return	nil,fmt.Errorf("invalid code")
	}

	dt := time.Now()
	item.UpdatedAt = dt
	u.setStatus(item)
	item.ID = id

	u.repo.ModifyItem(id,*item)

	return item, nil
}

func (u *itemUsecase) setStatus(item *dom.Item) {
	if item.Stock == 0 {
		item.Status = "INACTIVE"
	} else {
		item.Status = "ACTIVE"
	}
}

func (u *itemUsecase) GetItemsByStatusAndLimit(status string, limit int) ([]dom.Item,error){
	var db []dom.Item

	if status == "ACTIVE" || status == "INACTIVE" {
		db = u.repo.GetItemsByStatus(status)
	} else if status == "" {
		db = u.repo.GetDB()
	} else {
		return nil,fmt.Errorf("invalid url")
	}

	sort.Slice(db, func(i, j int) bool {
		return db[i].UpdatedAt.After(db[j].UpdatedAt)
	})

	if limit > len(db){
		limit = len(db)
	}

	return db[0:limit],nil
}