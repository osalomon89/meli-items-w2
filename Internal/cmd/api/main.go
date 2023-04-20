package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"meli-items-w2/Internal/adapters/controller"
	"meli-items-w2/Internal/adapters/repository"
	"meli-items-w2/Internal/usecase"
)

const port = ":9000"

func main() {

	r := gin.Default()

	itemRepository := repository.NewItemRepository()
	itemUseCase := usecase.NewItemUsecase(itemRepository)
	itemController := controller.NewItemContrller(itemUseCase)

	r.GET("/", controller.Index)
	r.POST("api/items", itemController.AddItem)
	r.PUT("api/items/:id", itemController.UpdateItem)
	r.GET("api/items", itemController.GetItem)
	r.DELETE("api/items/:id", itemController.DeleteItem)
	//r.GET("v1/items/:id", itemController.GetItems)

	r.Run(port)
	if err := r.Run(port); err != nil {
		log.Fatalln(err)
	}

}
