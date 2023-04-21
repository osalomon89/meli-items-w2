package main

import (
	"github.com/gin-gonic/gin"
	//dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"

	ctrl "github.com/javmoreno-meli/meli-item-w2/internal/infrastructure/controller"
	repository "github.com/javmoreno-meli/meli-item-w2/internal/infrastructure/repository"
	usecase "github.com/javmoreno-meli/meli-item-w2/internal/usecase"
)

// Puerto en el que correra nuestra API
const port = ":9001"

func main() {
	/* //creemos un item (articulo)
	itemOne := dom.Item{
		ID:          1,
		Code:        "JAV01",
		Title:       "Silla Ergonomica",
		Description: "Silla no solo para sentarse si no para sentarse bien :V",
		Price:       454900,
		Stock:       5,
		Status:      "ACTIVE",
	}
	itemTwo := dom.Item{
		ID:          2,
		Code:        "JAV02",
		Title:       "Escritorio de madera",
		Description: "Escritorio para oficina en madera",
		Price:       434900,
		Stock:       10,
		Status:      "ACTIVE",
	}
	itemThree := dom.Item{
		ID:          3,
		Code:        "JAV03",
		Title:       "Escritorio de metal",
		Description: "Escritorio para oficina en metal",
		Price:       600000,
		Stock:       0,
		Status:      "INACTTIVE",
	} */

	serviceItemRepository := repository.NewItemRepository()
	serviceUseCaseRepository := usecase.NewItemUseCase(serviceItemRepository)
	serviceItemController := ctrl.NewItemController(serviceUseCaseRepository)

	route := gin.Default()
	//Routes
	//listar todos los items en la base de datos (variable db)
	route.GET("/v1/items", serviceItemController.GetItems)
	//Guardar un item
	route.POST("/v1/items", serviceItemController.AddItems)
	//Listar item by ID
	route.GET("/v1/items/:id", serviceItemController.GetItemsById)
	//Actualizar item by ID
	//route.PUT("/v1/items/:id", serviceItemController.UpdateItems)
	//Eliminar item by Id
	//route.DELETE("/v1/items/:id", serviceItemController.DeleteItem)

	//Hagamos que nuestras Api corra en el puerto que definimos (9001)
	route.Run(port)
}

//Pendiente
//Agregar funcion que valide si hay otro codigo creado
/*

Refactorizar codigo -->

CREAR FUNCION PARA GUARDAR ITEM (simplifica additem) -
CREAR FUNCION PARA ACTUALIZAR ITEM (simplifica update)


*/
