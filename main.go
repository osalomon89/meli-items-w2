package main

import (
	"github.com/gin-gonic/gin"
	"log"
	ctlr "meli-items-w2/Internal/controller"
	"meli-items-w2/Internal/usecase"
)

const port = ":9000"

func main() {

	r := gin.Default()

	itemUseCase := usecase.NewItemUsecase()
	itemController := ctlr.NewItemContrller(itemUseCase)

	r.GET("/", ctlr.Index)
	r.GET("v1/items", itemController.GetItem)
	r.POST("v1/items", itemController.AddItem)
	r.GET("v1/items/:id", ctlr.GetItemById)
	r.PUT("v1/items/:id", ctlr.UpdateItem)
	r.DELETE("v1/items/:id", ctlr.DeleteItem)

	r.Run(port)
	if err := r.Run(port); err != nil {
		log.Fatalln(err)
	}
	/* log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	} */
}
