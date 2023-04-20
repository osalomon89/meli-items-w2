package repository

import (
	domain "apiRestPractice/app/domain/item_domain"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type itemRepository struct {
	items []domain.Item
}

var data []domain.Item

func DataMock() {
	jsonFile, errorJson := os.Open("assets/items.json")
	if errorJson != nil {
		fmt.Printf("errorJson: %s", errorJson.Error())
	}
	defer jsonFile.Close()

	byteValue, errorByte := io.ReadAll(jsonFile)
	if errorByte != nil {
		fmt.Printf("errorByte: %s", errorByte.Error())
	}

	var result []domain.Item
	json.Unmarshal(byteValue, &result)

	data = result
}

func NewItemRepository() domain.ItemRepository {
	return &itemRepository{}
}

func (i *itemRepository) AddItem(item domain.Item) error {

	fmt.Println("Guardando el item...")
	data = append(data, item)
	fmt.Println("Item guardado")
	return nil
}

func (i *itemRepository) UpdateItem(item domain.Item) error {
	var idFoundFlag = true
	for _, itemTemp := range data {
		if item.Id == itemTemp.Id {
			idFoundFlag = false
			fmt.Println("Modificando el item...")
			data[item.Id] = item
			return nil
		}
	}
	if idFoundFlag {
		return fmt.Errorf("id not found")
	}
	return nil
}

func (i *itemRepository) GetById(id int) (*domain.Item, error) {
	var idFoundFlag = true
	for _, itemTemp := range data {
		if id == itemTemp.Id {
			idFoundFlag = false
			fmt.Println(" Item encontrado")
			return &data[id], nil
		}
	}
	if idFoundFlag {
		return nil, fmt.Errorf("id not found")
	}
	return nil, nil
}

func (i *itemRepository) DeleteById(id int) error {
	var idFoundFlag = true
	for index, itemTemp := range data {
		if id == itemTemp.Id {
			idFoundFlag = false
			data = append(data[:index], data[index+1:]...)
			fmt.Println(" Item eliminado")
			return nil
		}
	}
	if idFoundFlag {
		return fmt.Errorf("id not found")
	}
	return nil
}

func (i *itemRepository) GetByStatusAndLimit(status string, limit int) []domain.Item {
	var itemsRes []domain.Item
	if status == "ACTIVE" || status == "INACTIVE" {
		for _, item := range data {
			if status == item.Status && len(itemsRes) < limit {
				itemsRes = append(itemsRes, item)
			}
		}
	} else if status == "" {
		for _, item := range data {
			if len(itemsRes) < limit {
				itemsRes = append(itemsRes, item)
			}
		}
	}
	return itemsRes
}
