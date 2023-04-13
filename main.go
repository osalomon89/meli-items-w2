package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const port = ":9000"

type Item struct {
	Id          int     `json:"id"`
	Code        string  `json:"code" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Descripcion string  `json:"descripcion" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Status      string  `json:"status"`
	CreatAt     string  `json:"creat_at"`
	UpdateAt    string  `json:"update_at"`
	Author      string  `json:"author"`
}

var db []Item

func main() {

	item1 := Item{
		Id:          rand.Int(),
		Code:        "escritorio1234",
		Title:       "Escritorio",
		Descripcion: "Excelente escritorio confortable",
		Price:       24500,
		Stock:       10,
		Status:      "Activo",
		CreatAt:     "2020-05-10T04:20:33Z",
		UpdateAt:    "2020-05-10T05:30:00Z",
	}

	item2 := Item{
		Id:          rand.Int(),
		Code:        "Sofa1234",
		Title:       "Sofa",
		Descripcion: "Comodo sofa para tardear",
		Price:       24500,
		Stock:       10,
		Status:      "Activo",
		CreatAt:     "2020-05-10T04:20:33Z",
		UpdateAt:    "2020-05-10T05:30:00Z",
	}

	item3 := Item{
		Id:          rand.Int(),
		Code:        "Iphone1234",
		Title:       "Iphone 20",
		Descripcion: "Dispositivo de alta gama",
		Price:       45687,
		Stock:       7,
		Status:      "Activo",
		CreatAt:     "2020-05-10T04:20:33Z",
		UpdateAt:    "2020-05-10T05:30:00Z",
	}

	db = append(db, item1, item2, item3)

	r := gin.Default()

	r.GET("/", index)
	r.GET("v1/items", getItem)
	r.POST("v1/items", addItem)
	r.GET("v1/items/:id", getItemById)
	r.PUT("v1/items/:id", updateItem)
	r.DELETE("v1/items/:id", deleteItem)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

// Función que permite agregar items
func addItem(c *gin.Context) {

	request := c.Request
	var item Item
	err := json.NewDecoder(request.Body).Decode(&item)
	if err := validate.Struct(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}
	if VerificaRepetido(item) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Código repetido", err.Error()),
		})
		return

	}

	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}

	db = append(db, item)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

// Funcion para obtener items
func getItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return

	}
	var item Item

	for i, v := range db {
		if v.Id == id {
			db[i] = item
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "Item no encontrado",
	})

}

// Función que permite acutalizar items
func updateItem(c *gin.Context) {
	r := c.Request
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	var item Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err := validate.Struct(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}
	if VerificaRepetido(item) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Código repetido", err.Error()),
		})
		return

	}

	for i, v := range db {
		if v.Id == id {
			db[i] = item
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

// Función que permite obtener items dado un id
func getItemById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	for _, v := range db {
		if v.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"error": false,
				"data":  v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "Item no encontrado",
	})
}

// Función que permite elimar items
func deleteItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	for i, v := range db {
		if v.Id == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "Item no encontrado",
	})
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

// Función que nos dice si el código es repetivo
func VerificaRepetido(item Item) bool {
	for _, i := range db {
		if i.Code == item.Code {
			return true
		}
	}
	return false
}
