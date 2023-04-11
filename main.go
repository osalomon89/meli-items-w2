package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const port = ":9000"

type Item struct {
	ID          int     `json:"id"`
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Descripcion string  `json:"descripcion"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
	CreatAt     string  `json:"creat_at"`
	UpdateAt    string  `json:"update_at"`
	Author      string  `json:"author"`
}

var db []Item

func main() {
	item1 := Item{
		ID:          1,
		Code:        "asiu654",
		Title:       "Escritorio",
		Descripcion: "Excelente escritorio confortable",
		Price:       24500,
		Stock:       10,
	}

	item2 := Item{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	item3 := Item{
		ID:          3,
		Code:        "abvc887",
		Title:       "Lavadora",
		Descripcion: "Excelente lavadora",
		Price:       500,
		Stock:       5,
	}

	db = append(db, item1, item2, item3)

	r := gin.Default()

	r.GET("/", index)
	r.GET("/books", getItem)
	r.POST("/books", addItem)
	r.GET("/books/:id", getItemById)
	r.PUT("/books/:id", updateItem)
	r.DELETE("/books/:id", deleteItem)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

// Actualizando item
func updateItem(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	var item Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range db {
		if v.ID == id {
			db[i] = item
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

// Agregar item/ ///////
func addItem(ctx *gin.Context) {

	request := ctx.Request
	var item Item
	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	db = append(db, item)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

func getItem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

func getItemById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"error": false,
				"data":  v,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "book not found",
	})
}

func deleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}
