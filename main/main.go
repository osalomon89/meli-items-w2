package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"

	dom "meli-items-w2/internal/domain"
	"meli-items-w2/internal/infraestructure/controller"
	"meli-items-w2/internal/infraestructure/repository"
	"meli-items-w2/internal/usecase"
)

// Puerto de escucha declarado como const
const port string = "localhost:8888"

func main() {

	// Router default de gin y logueo
	router := gin.Default()

	serviceRepository := repository.NewItemRepository()
	serviceUsecase := usecase.NewItemUsecase(serviceRepository)
	serviceController := controller.NewItemController(serviceUsecase)

	// Instancio 3 items para agregar a la BD
	var item1 = dom.Item{
		Id:          1,
		Code:        "SAM27324354",
		Title:       "Tablet Samsung Galaxy Tab S7",
		Description: "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       550000,
		Stock:       3,
		Status:      "ACTIVE",
		CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		UpdatedAt:   time.Time{},
	}
	var item2 = dom.Item{
		Id:          2,
		Code:        "SAM27324355",
		Title:       "Tablet Samsung Galaxy Tab S8",
		Description: "Galaxy Tab S8 with S Pen SM-t733 12.4 pulgadas y 8GB de memoria RAM",
		Price:       950000,
		Stock:       2,
		Status:      "ACTIVE",
		CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		UpdatedAt:   time.Time{},
	}
	var item3 = dom.Item{
		Id:          3,
		Code:        "SAM27324356",
		Title:       "Smartphone Samsung J2",
		Description: "Smarthphone J2 6 pulgadas y 1GB de memoria RAM",
		Price:       300000,
		Stock:       42,
		Status:      "ACTIVE",
		CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		UpdatedAt:   time.Time{},
	}

	// Agrego los items a la DB
	serviceRepository.AddItem(item1)
	serviceRepository.AddItem(item2)
	serviceRepository.AddItem(item3)

	// ******** ENDPOINTS *******

	// Get
	router.GET("v1/items", serviceController.ListItem)
	router.GET("v1/items/:id", serviceController.GetItemByID)

	// Post
	router.POST("v1/items", serviceController.AddItem)

	// Put
	router.PUT("v1/items/:id", serviceController.UpdateItem)

	// Delete
	router.DELETE("v1/items/:id", serviceController.DeleteItem)

	// Prendemos la maquinola
	router.Run(port)

	// Mensaje del puerto
	log.Println("server listening to the port:", port)
}
