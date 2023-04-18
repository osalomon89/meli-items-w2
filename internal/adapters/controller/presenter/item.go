package presenter

import (
	"time"

	"github.com/osalomon89/meli-items-w2/internal/entity"
)

type ItemJson struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Item(i entity.Item) ItemJson {
	var responseItem ItemJson

	responseItem.Id = i.Id
	responseItem.Code = i.Code
	responseItem.Title = i.Title
	responseItem.Description = i.Description
	responseItem.Price = i.Price
	responseItem.Stock = i.Stock
	responseItem.Status = i.Status
	responseItem.CreatedAt = i.CreatedAt
	responseItem.UpdatedAt = i.UpdatedAt

	return responseItem
}

func Items(i []entity.Item) []ItemJson {
	var responseItems []ItemJson
	for _, v := range i {
		responseItems = append(responseItems, Item(v))
	}

	return responseItems
}
