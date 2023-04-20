package main

import (
	"apiRestPractice/app/adapters/item_adapters/controller"
	"apiRestPractice/app/adapters/item_adapters/repository"
	"apiRestPractice/app/usecase"
	"github.com/gin-gonic/gin"
)

const port = ":9000"

func main() {
	repository.DataMock()
	server := gin.Default()

	repo := repository.NewItemRepository()
	itemUsecase := usecase.NewItemUsecase(repo)
	ctrl := controller.NewItemController(itemUsecase)

	server.POST("v1/items", ctrl.AddItem)
	server.PUT("v1/items/:id", ctrl.UpdateItem)
	server.GET("v1/items/:id", ctrl.GetById)
	server.DELETE("v1/items/:id", ctrl.DeleteById)
	server.GET("v1/items", ctrl.GetByStatusAndLimit)
	if err := server.Run(port); err != nil {
		panic(err)
	}
}
