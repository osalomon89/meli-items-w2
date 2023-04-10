package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
)

const port = ":9000"

type Item struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var db []Item

func main() {
	fmt.Println(db)

	r := gin.Default()

	r.POST("v1/items", addItem)
	r.PUT("v1/items/:id", updateItem)
	r.GET("v1/items/:id", getItem)
	r.DELETE("v1/items/:id", deleteItem) 
	r.GET("v1/items", getItems) // como expresar el opcional en visual

	r.Run(port)

}

func addItem(c *gin.Context) {
	request := c.Request
	body := request.Body

	var item Item
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	// Chequeo de code unico
	// Falta personalizar el error
	if codeRepetido(&item) {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	item.ID = obtenerSiguienteID()
	setStatus(&item)
	dt := time.Now()
	item.CreatedAt = dt.String()
	item.UpdatedAt = dt.String()

	db = append(db, item)

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}

// Funciones para los endpoints

func updateItem(c *gin.Context) {
	request := c.Request
	body := request.Body

	idParam := c.Param("id")
	id, err1 := strconv.Atoi(idParam)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err1.Error()),
		})
		return
	}

	var item Item
	err2 := json.NewDecoder(body).Decode(&item)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err2.Error()),
		})
		return
	}

	// Chequeo de code unico
	// Falta personalizar el error
	if codeRepetido(&item) {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err2.Error()),
		})
		return
	}

	setStatus(&item)
	dt := time.Now()
	for _, val := range db {
		if val.ID == id {
			val.Code = item.Code
			val.Title = item.Title
			val.Description = item.Description
			val.Price = item.Price
			val.Stock = item.Stock
			val.Status = item.Status
			val.UpdatedAt = dt.String()

			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  val,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  "item not found",
	})
}

func getItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  "item not found",
	})

}

func deleteItem(c *gin.Context) {
	// FALTA IMPLEMENTAR
}

func getItems(c *gin.Context){
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  db,
	})
}




// Funciones auxiliares

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Devuelve true si el code ya existe
func codeRepetido(item *Item) bool {
	var repetido bool
	for _, val := range db {
		if val.Code == item.Code {
			repetido = true
		}
	}
	return repetido
}

// Obtiene el próximo ID libre
func obtenerSiguienteID() int {
	var idSiguiente int
	for _, val := range db {
		if idSiguiente < val.ID {
			idSiguiente = val.ID
		}
	}
	idSiguiente++
	return idSiguiente
}

// Setea el status en funcion del stock
func setStatus(item *Item) {
	if item.Stock == 0 {
		item.Status = "INACTIVE"
	} else {
		item.Status = "ACTIVE"
	}
}
