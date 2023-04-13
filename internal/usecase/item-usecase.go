package usecase

import (
	"fmt"
	"time"

	dom "github.com/osalomon89/meli-items-w2/internal/domain"
)

type ItemUsecase interface {
	GetItemByID(id int) *dom.Item
	DeleteItemByID(id int) bool
	GetAllItems() []dom.Item
	AddItemByItem(item *dom.Item) (*dom.Item,error)
	setStatus(item *dom.Item)
	UpdateItemByItem(id int,item *dom.Item) (*dom.Item,error)
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
	if u.repo.CodeRepetido(*item) {
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
	if u.repo.CodeRepetido(*item) {
		return	nil,fmt.Errorf("invalid code")
	}

	dt := time.Now()
	item.UpdatedAt = dt
	u.setStatus(item)

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