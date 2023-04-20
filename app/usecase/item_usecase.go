package usecase

import (
	domain "apiRestPractice/app/domain/item_domain"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ItemUsecase interface {
	AddItem(item map[string]any) error
	UpdateItem(item map[string]any, id int) error
	GetById(id int) (*domain.Item, error)
	DeleteById(id int) error
	GetByStatusAndLimit(status string, limit int) ([]domain.Item, error)
}

type itemUsecase struct {
	repo domain.ItemRepository
}

func NewItemUsecase(repo domain.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (i itemUsecase) AddItem(item map[string]any) error {
	if len(checkFields(item)) > 0 {
		var errToString string
		for _, error := range checkFields(item) {
			errToString = errToString + fmt.Sprintf("%s, ", error.Error())
		}
		return fmt.Errorf(errToString)
	}
	return i.repo.AddItem(newItem(item))
}

func (i itemUsecase) UpdateItem(item map[string]any, id int) error {
	if len(checkFields(item)) > 0 {
		var errToString string
		for _, error := range checkFields(item) {
			errToString = errToString + fmt.Sprintf("%s, ", error.Error())
		}
		return fmt.Errorf(errToString)
	}
	itemToUpdate, errorGet := i.repo.GetById(id)
	if errorGet != nil {
		return errorGet
	}
	return i.repo.UpdateItem(updateItem(item, *itemToUpdate))
}

func (i itemUsecase) GetById(id int) (*domain.Item, error) {
	itemToGet, errorGet := i.repo.GetById(id)
	if errorGet != nil {
		return nil, errorGet
	}
	return itemToGet, nil
}

func (i itemUsecase) DeleteById(id int) error {
	return i.repo.DeleteById(id)
}

func (i itemUsecase) GetByStatusAndLimit(status string, limit int) ([]domain.Item, error) {

	if limit == 0 {
		limit = 10
	}
	if limit > 20 {
		return nil, fmt.Errorf("the maximum limit value is 20")
	}
	if status == "ACTIVE" || status == "INACTIVE" || status == "" {
		return i.repo.GetByStatusAndLimit(status, limit), nil
	} else {
		return nil, fmt.Errorf("invalid status")
	}
}

func checkFields(req map[string]any) []error {
	var errors []error
	if req["code"] == "" || req["code"] == nil {
		errors = append(errors, fmt.Errorf("code is required"))
	}
	if req["title"] == "" || req["title"] == nil {
		errors = append(errors, fmt.Errorf("title is required"))
	}
	if req["description"] == "" || req["description"] == nil {
		errors = append(errors, fmt.Errorf("description is required"))
	}
	if req["price"] == "" || req["price"] == nil {
		errors = append(errors, fmt.Errorf("price is required"))
	}
	if req["stock"] == "" || req["stock"] == nil {
		errors = append(errors, fmt.Errorf("stock is required"))
	}
	return errors
}

func newItem(result map[string]any) (item domain.Item) {
	location, errLo := time.LoadLocation("Africa/Conakry")
	if errLo != nil {
		fmt.Printf("Error in location: %s", errLo.Error())
	}
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	item.Code = fmt.Sprint(result["code"])
	item.Title = fmt.Sprint(result["title"])
	item.Description = fmt.Sprint(result["description"])
	item.Price, _ = strconv.Atoi(fmt.Sprint(result["price"]))
	item.Stock, _ = strconv.Atoi(fmt.Sprint(result["stock"]))
	dateStrg := strings.Split(currentTime, " ")
	item.CreatedAt = dateStrg[0] + "T" + dateStrg[1] + "Z"
	item.UpdatedAt = item.CreatedAt
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
	return item
}

func updateItem(result map[string]any, itemToUpdate domain.Item) (item domain.Item) {
	location, errLo := time.LoadLocation("Africa/Conakry")
	if errLo != nil {
		fmt.Printf("Error in location: %s", errLo.Error())
	}
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	item.Code = fmt.Sprint(result["code"])
	item.Title = fmt.Sprint(result["title"])
	item.Description = fmt.Sprint(result["description"])
	item.Price, _ = strconv.Atoi(fmt.Sprint(result["price"]))
	item.Stock, _ = strconv.Atoi(fmt.Sprint(result["stock"]))
	item.Id = itemToUpdate.Id
	dateStrg := strings.Split(currentTime, " ")
	item.CreatedAt = itemToUpdate.CreatedAt
	item.UpdatedAt = dateStrg[0] + "T" + dateStrg[1] + "Z"
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
	return item
}
