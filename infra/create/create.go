package create

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/NitroML/meli-items-w2/domain"
	"github.com/NitroML/meli-items-w2/infra/getall"
)

var (
	data []domain.Item
	once sync.Once
)

func ItemCreate(item domain.Item) (domain.Item, error) {
	// Load all items
	items, err := getall.ItemsGetAll()
	if err != nil {
		return domain.Item{}, err
	}

	// Generate new ID
	newID := 1
	if len(items) > 0 {
		lastItem := items[len(items)-1]
		newID = lastItem.ID + 1
	}

	// Set new ID and timestamps for item
	item.ID = newID
	item.Created = time.Now()
	item.Updated = time.Now()

	// Add new item to data
	data = append(data, item)

	// Write data back to JSON file
	filePath, err := filepath.Abs("infra/db.json")
	if err != nil {
		return domain.Item{}, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return domain.Item{}, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(map[string]interface{}{"items": data})
	if err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func saveToFile() error {
	filePath, err := filepath.Abs("infra/db.json")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
