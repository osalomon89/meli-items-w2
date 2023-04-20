package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "meli-items-w2/Internal/adapters/controller"
)

const port = ":9000"

func NewHTTPServer(bookCtrl ctrl.BookController) error {
	r := gin.Default()

	basePath := "/api/v1/inventory"
	publicRouter := r.Group(basePath)

	publicRouter.GET("/books", bookCtrl.GetBooks)
	publicRouter.POST("/books", bookCtrl.AddBook)
	publicRouter.GET("/books/:id", bookCtrl.GetBook)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, r)
}
