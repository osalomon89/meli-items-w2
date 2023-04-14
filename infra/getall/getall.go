package getall

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/NitroML/meli-items-w2/domain"
)

var (
	data []domain.Item
	mu   sync.Mutex
	once sync.Once
)

func ItemsGetAll() ([]domain.Item, error) {
	once.Do(func() {
		filePath, err := filepath.Abs("infra/db.json")
		if err != nil {
			panic(err)
		}

		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
	})

	mu.Lock()
	defer mu.Unlock()

	// Return a copy of the data slice to prevent concurrent access issues
	return append([]domain.Item(nil), data...), nil
}
