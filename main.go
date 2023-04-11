package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Puerto de escucha declarado como const
const port string = "localhost:8888"

// Creo BD local
var itemsDB []Item

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

func main() {
	// Instancio 3 items para agregar a la BD
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
		},

		{
			Id:          3,
			Code:        "SAM27324356",
			Title:       "Smartphone Samsung J2",
			Description: "Smarthphone J2 6 pulgadas y 1GB de memoria RAM",
			Price:       300000,
			Stock:       42,
			Status:      "ACTIVE",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Agrego los items a la DB
	itemsDB = append(items)
	fmt.Println(itemsDB)

	// Router default de gin y logueo
	router := gin.Default()

	// ******** ENDPOINTS *******

	router.GET("v1/items", listItems)

	/*
		server.POST("v1/items", addItem)
		server.PUT("v1/items/:id", updateItem)
		server.GET("v1/items/:id", getItemByID)
		server.DELETE("v1/items/:id", deleteItem)
	*/

	// Prendemos la maquinola
	router.Run(port)

	// Mensaje del puerto
	log.Println("server listening to the port:", port)

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
	if len(itemsDB) == 0 {
		c.Error(fmt.Errorf("No hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, itemsDB)
}

// Modificar item
func updateItem() {

}

// Eliminar item
func deleteItem() {

}
