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

	//r.GET("/", ctlr.Index)
	r.POST("v1/items", itemController.AddItem)
	r.PUT("v1/items/:id", itemController.UpdateItem)
	r.GET("v1/items", itemController.GetItem)
	r.DELETE("v1/items/:id", itemController.DeleteItem)
	//r.GET("v1/items/:id", itemController.GetItems)

	r.Run(port)
	if err := r.Run(port); err != nil {
		log.Fatalln(err)
	}
	/* log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	} */
}
