package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/osalomon89/meli-items-w2/internal/infrastructure/controller"
	"github.com/osalomon89/meli-items-w2/internal/infrastructure/repository"
	"github.com/osalomon89/meli-items-w2/internal/usecase"
)

const port = ":8080"

func main() {

	server := gin.Default()

	itemRepository := repository.NewItemRepository()
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemController := controller.NewItemController(itemUsecase)

	server.POST("v1/items", itemController.AddItem)
	server.PUT("v1/items/:id", itemController.UpdateItem)
	server.GET("v1/items/:id", itemController.GetItem)
	server.DELETE("v1/items/:id", itemController.DeleteItem)
	server.GET("v1/items", itemController.GetItems)

	server.Run(port)

	if err := server.Run(port); err != nil {
		log.Fatalln(err)
	}

}
