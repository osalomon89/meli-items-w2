package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osalomon89/meli-items-w2/internal/adapters/controller"
)

const port = ":8080"

func NewHTTPServer(itemController controller.ItemController) error {

	server := gin.Default()

	basePasth := "v1"
	publicRouter := server.Group(basePasth)

	publicRouter.POST("/items", itemController.AddItem)
	publicRouter.PUT("/items/:id", itemController.UpdateItem)
	publicRouter.GET("/items/:id", itemController.GetItem)
	publicRouter.DELETE("/items/:id", itemController.DeleteItem)
	publicRouter.GET("/items", itemController.GetItems)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, server)
}
