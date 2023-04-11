package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
			CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
			UpdatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		},

		{
			Id:          2,
			Code:        "SAM27324355",
			Title:       "Tablet Samsung Galaxy Tab S8",
			Description: "Galaxy Tab S8 with S Pen SM-t733 12.4 pulgadas y 8GB de memoria RAM",
			Price:       950000,
			Stock:       2,
			Status:      "ACTIVE",
			CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
			UpdatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		},

		{
			Id:          3,
			Code:        "SAM27324356",
			Title:       "Smartphone Samsung J2",
			Description: "Smarthphone J2 6 pulgadas y 1GB de memoria RAM",
			Price:       300000,
			Stock:       42,
			Status:      "ACTIVE",
			CreatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
			UpdatedAt:   time.Date(2023, 4, 11, 17, 0, 0, 0, time.FixedZone("-05:00", -5*60*60)),
		},
	}

	// Agrego los items a la DB
	saveItem(items)

	// Router default de gin y logueo
	router := gin.Default()

	// ******** ENDPOINTS *******

	// Get
	router.GET("v1/items", listItems)
	router.GET("v1/items/:id", getItemByID)

	// Post
	router.POST("v1/items", addItem)
	/* Para probar
	[
	    {
	        "code": "LPTP27324354",
	        "title": "Laptop Dell XPS 13",
	        "description": "Dell XPS 13 con procesador Intel Core i7 de 11.ª generación y 16GB de memoria RAM",
	        "price": 2000000,
	        "stock": 5,
	        "status": "ACTIVE"
	    },
	    {
	        "code": "LPTP27324355",
	        "title": "Laptop HP Spectre x360",
	        "description": "HP Spectre x360 con procesador Intel Core i7 de 11.ª generación y 8GB de memoria RAM",
	        "price": 1800000,
	        "stock": 3,
	        "status": "ACTIVE"
	    }
	]
	*/

	// Put
	router.PUT("v1/items/:id", updateItem)

	// Delete
	router.DELETE("v1/items/:id", deleteItem)

	// Prendemos la maquinola
	router.Run(port)

	// Mensaje del puerto
	log.Println("server listening to the port:", port)

}

// Guardar un item
func saveItem(addItem []Item) {
	itemsDB = append(itemsDB, addItem...)
}

// Añadir item
func addItem(c *gin.Context) {
	var newSliceItem []Item
	c.BindJSON(&newSliceItem)

	for i := range newSliceItem {
		newSliceItem[i].Id = len(itemsDB) + i + 1
		newSliceItem[i].CreatedAt = time.Now()
		newSliceItem[i].UpdatedAt = time.Now()
	}

	saveItem(newSliceItem)
	c.IndentedJSON(http.StatusCreated, newSliceItem)

}

// Obtener Item por id
func getItemByID(c *gin.Context) {
	// Obtener el ID del parámetro de la URL
	id := c.Param("id")

	// Variable bandera para verificar si sí se encuentra el id solicitado
	//isFound := false

	// Para guardar el item si se encuentra
	var itemFound Item

	// Casteando el param que llega en string a int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	// Buscamos el id que necesitamos
	for _, item := range itemsDB {
		if item.Id == idInt {
			itemFound = item
			break
		}
	}

	// Retornamos ok si encontramos el id, no es necesario igular a true en la condición ya que en Go porque el valor booleano es en sí una condición en el caso if found {...}
	// Se puede evitar el uso de la variable bandera si directamente preguntamos por el valor por default de la struct Item
	if itemFound != (Item{}) {
		c.JSON(http.StatusOK, itemFound)
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
	}

}

// Listar todos los items
func listItems(c *gin.Context) {
	if len(itemsDB) == 0 {
		c.Error(fmt.Errorf("no hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, itemsDB)
}

// Modificar item
func updateItem(c *gin.Context) {
	//var itemFound Item
	var itemToUpdate Item
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.BindJSON(&itemToUpdate)

	for _, item := range itemsDB {
		if item.Id == idInt {
			item.Code = itemToUpdate.Code
			item.Title = itemToUpdate.Title
			item.Description = itemToUpdate.Description
			item.Price = itemToUpdate.Price
			item.Stock = itemToUpdate.Stock
			item.Status = itemToUpdate.Status
			item.CreatedAt = item.CreatedAt
			item.UpdatedAt = time.Now()
			c.JSON(http.StatusOK, itemToUpdate)
			break
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, "No se encuentra el id solicitado.")
			break
		}
	}

}

// Eliminar item
func deleteItem(c *gin.Context) {

}
