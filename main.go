package main

import (
	"apiRestPractice/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	controller.DataMock()
	server := gin.Default()
	server.POST("v1/items", controller.AddItem)
	server.PUT("v1/items/:id", controller.UpdateItem)
	server.GET("v1/items/:id", controller.GetById)
	server.DELETE("v1/items/:id", controller.DeleteById)
	server.GET("v1/items", controller.GetByStatusAndLimit)
	server.Run(":9000")
}
