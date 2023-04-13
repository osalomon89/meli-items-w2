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
			UpdatedAt:   time.Time{},
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
			UpdatedAt:   time.Time{},
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
			UpdatedAt:   time.Time{},
		},
	}

	// Agrego los items a la DB
	saveItem(items)

	// Router default de gin y logueo
	router := gin.Default()

	// ******** ENDPOINTS *******

	// Get
	//router.GET("v1/items", listItems)
	router.GET("v1/items", getAllFiltered)
	router.GET("v1/items/:id", getItemByID)

	// Post
	router.POST("v1/items", addItem)

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

	// Manejando el error
	if err := c.BindJSON(&newSliceItem); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	for i := range newSliceItem {
		newSliceItem[i].Id = len(itemsDB) + i + 1
		newSliceItem[i].CreatedAt = time.Now()
		newSliceItem[i].UpdatedAt = time.Time{}
		if newSliceItem[i].Stock > 0 {
			newSliceItem[i].Status = "ACTIVE"
		} else {
			newSliceItem[i].Status = "INACTIVE"
		}

	}

	saveItem(newSliceItem)
	c.IndentedJSON(http.StatusCreated, newSliceItem)

}

// Obtener Item por id
func getItemByID(c *gin.Context) {
	// Obtener el ID del parámetro de la URL
	id := c.Param("id")

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

// obtener items con filtros

func getAllFiltered(c *gin.Context) {
	status := c.Query("status")

	var dbFiltered []Item

	if status != "ACTIVE" && status != "INACTIVE" && status != "ALL" {
		c.Error(fmt.Errorf("status inválido"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, item := range itemsDB {
		if item.Status == status {
			dbFiltered = append(dbFiltered, item)
		}

	}

	if len(itemsDB) == 0 {
		c.Error(fmt.Errorf("no hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if len(dbFiltered) == 0 {
		c.Error(fmt.Errorf("no hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, dbFiltered)

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

// Modificar item no funca
func updateItem(c *gin.Context) {
	//var itemFound Item
	var itemToUpdate Item
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if error := c.BindJSON(&itemToUpdate); error != nil {
		c.JSON(http.StatusNotFound, error)
		return
	}

	var itemToUpdatePtr *Item

	for i := range itemsDB {
		if itemsDB[i].Id == idInt {
			itemToUpdatePtr = &itemsDB[i]
			break
		}
	}

	if itemToUpdatePtr == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
		return
	}

	if itemToUpdate.Stock > 0 {
		itemToUpdate.Status = "ACTIVE"
	} else {
		itemToUpdate.Status = "INACTIVE"
	}

	itemToUpdatePtr.Code = itemToUpdate.Code
	itemToUpdatePtr.Title = itemToUpdate.Title
	itemToUpdatePtr.Description = itemToUpdate.Description
	itemToUpdatePtr.Price = itemToUpdate.Price
	itemToUpdatePtr.Stock = itemToUpdate.Stock
	itemToUpdatePtr.Status = itemToUpdate.Status
	itemToUpdatePtr.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, itemToUpdatePtr)

}

// Eliminar item
func deleteItem(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	for i, item := range itemsDB {
		if item.Id == idInt {
			itemsDB = append(itemsDB[:i], itemsDB[i+1:]...)
			c.JSON(http.StatusOK, "item eliminado con éxito id: "+id)
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
}
