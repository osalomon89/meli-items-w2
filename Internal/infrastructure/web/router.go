package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "meli-items-w2/Internal/adapters/controller"
)

const port = ":9000"

func NewHTTPServer(itemCtrl ctrl.ItemController) error {
	r := gin.Default()

	basePath := "/api/v1/inventory"
	publicRouter := r.Group(basePath)

	publicRouter.GET("/items", itemCtrl.GetItem)
	publicRouter.POST("/items", itemCtrl.AddItem)
	publicRouter.GET("/items/:id", itemCtrl.UpdateItem)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, r)
}
