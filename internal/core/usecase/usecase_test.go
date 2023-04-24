package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"meli-items-w2/internal/core/usecase"
	"meli-items-w2/internal/mocks"
)

// Definimos un mock para nuestro repositorio que nos permitirá simular el comportamiento de la base de datos
type itemRepositoryMock struct {
	err error
}

// Implementamos el método GetItemByID del repositorio para retornar el error que definimos en la estructura
func (repo itemRepositoryMock) GetItemByID(id int) error {
	return repo.err
}

func Test_itemUsecase_GetItemById(t *testing.T) {
	assert := assert.New(t)

	// Definimos los argumentos de entrada del método que vamos a testear
	type args struct {
		id int
	}

	// Definimos los casos de prueba que vamos a testear
	tests := []struct {
		name      string
		args      args
		repoError error
		repoTimes int
		wantedErr error
	}{
		{
			name:      "Should work correctly",
			wantedErr: nil,
			args: args{
				id: 1,
			},
			repoError: nil,
			repoTimes: 1,
		},
		{
			name:      "Should return error id doesn't exist",
			wantedErr: fmt.Errorf("item id does not exist"),
			args: args{
				id: 100,
			},
			repoError: nil,
			repoTimes: 0,
		},
		{
			name:      "Should return error when repository returns an error",
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
			args: args{
				id: 1,
			},
			repoError: errors.New("the repository error"),
			repoTimes: 1,
		},
	}

	// Recorremos todos los casos de prueba y los ejecutamos
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Creamos un nuevo controlador de mocks para cada caso de prueba
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Creamos un mock del repositorio que vamos a utilizar
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			// Indicamos cuál va a ser el comportamiento esperado del mock
			repositoryMock.EXPECT().
				GetItemByID(tt.args.id).
				Return(tt.repoError).
				Times(tt.repoTimes)

			// Creamos una nueva instancia del servicio que vamos a testear,
			// pasándole el mock creado anteriormente
			svc := usecase.NewItemUsecase(repositoryMock)

			// Ejecutamos el método que queremos testear
			err := svc.GetItemById(tt.args.id)

			// Verificamos que la respuesta obtenida es la que esperábamos
			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}

			assert.Nil(err)
		})
	}

}

/*
import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"meli-items-w2/internal/core/usecase"
	"testing"
)

// Definimos un mock para nuestro repositorio que nos permitirá simular el comportamiento de la base de datos
type itemRepositoryMock struct {
	err error
}

// Implementamos el método SaveItem del repositorio para retornar el error que definimos en la estructura
func (repo itemRepositoryMock) SaveItem(name string, stock int) error {
	return repo.err
}

// Implementamos el método GetItemByID del repositorio para retornar el error que definimos en la estructura
func (repo itemRepositoryMock) GetItemByID(itemID uint) error {
	return repo.err
}

func Test_itemService_CreateItem(t *testing.T) {
	assert := assert.New(t)

	// Definimos los argumentos de entrada del método que vamos a testear
	type args struct {
		name  string
		stock int
	}

	// Definimos los casos de prueba que vamos a testear
	tests := []struct {
		name      string
		args      args
		repoError error
		repoTimes int
		wantedErr error
	}{
		{
			name:      "Should work correctly",
			wantedErr: nil,
			args: args{
				name:  "tablet",
				stock: 10,
			},
			repoError: nil,
			repoTimes: 1,
		},
		{
			name:      "Should return error when item name is empty",
			wantedErr: fmt.Errorf("item name could not be empty"),
			args: args{
				name:  "",
				stock: 10,
			},
			repoError: nil,
			repoTimes: 0,
		},
		{
			name:      "Should return error when item stock is zero",
			wantedErr: fmt.Errorf("stock could not be zero"),
			args: args{
				name:  "tablet",
				stock: 0,
			},
			repoError: nil,
			repoTimes: 0,
		},
		{
			name:      "Should return error when repository returns an error",
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
			args: args{
				name:  "tablet",
				stock: 10,
			},
			repoError: errors.New("the repository error"),
			repoTimes: 1,
		},
	}

	// Recorremos todos los casos de prueba y los ejecutamos
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Creamos un nuevo controlador de mocks para cada caso de prueba
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Creamos un mock del repositorio que vamos a utilizar
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			// Indicamos cuál va a ser el comportamiento esperado del mock
			repositoryMock.EXPECT().
				SaveItem(tt.args.name, tt.args.stock).
				Return(tt.repoError).
				Times(tt.repoTimes)

			// Creamos una nueva instancia del servicio que vamos a testear,
			// pasándole el mock creado anteriormente
			svc := usecase.NewItemUsecase(repositoryMock)

			// Ejecutamos el método que queremos testear
			err := svc.CreateItem(tt.args.name, tt.args.stock)

			// Verificamos que la respuesta obtenida es la que esperábamos
			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}

			assert.Nil(err)
		})
	}
}*/
