package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var db []Item

const port string = "localhost:8888"

func main() {
	var items = []Item{
		{
			Id:          1,
			Code:        "SAM27324354",
			Title:       "Tablet Samsung Galaxy Tab S7",
			Description: "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
			Price:       550000,
			Stock:       3,
			Status:      "ACTIVE",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},

		{
			Id:          2,
			Code:        "SAM27324355",
			Title:       "Tablet Samsung Galaxy Tab S8",
			Description: "Galaxy Tab S8 with S Pen SM-t733 12.4 pulgadas y 8GB de memoria RAM",
			Price:       950000,
			Stock:       2,
			Status:      "ACTIVE",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}}
	fmt.Println(items)

	router := gin.Default()

	router.GET("v1/items", listItems)
	/*
		server.POST("v1/items", addItem)
		server.PUT("v1/items/:id", updateItem)
		server.GET("v1/items/:id", getItemByID)
		server.DELETE("v1/items/:id", deleteItem)
	*/

	router.Run(port)

}

// Item Creamos la estructura Item y las etiquetas del JSON
type Item struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ResponseInfo Creamos la estructura ResponseInfo y las etiquetas del JSON
type ResponseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Guardar un item
func saveItem(c *gin.Context) {

}

// Añadir item
func addItem() {

}

// Obtener Item por id
func getItemByID() {

}

// Listar todos los items
func listItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db)
}

// Modificar item
func updateItem() {

}

// Eliminar item
func deleteItem() {

}
