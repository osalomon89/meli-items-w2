package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	dom "github.com/osalomon89/neocamp-meli-w2/domain"
	controllerService "github.com/osalomon89/neocamp-meli-w2/controller"
	
	useCaseService "github.com/osalomon89/neocamp-meli-w2/usecases"
	
)

const port = ":9000"
/*
type Item struct {
	ID     int    `json:"id"`
	Code string `json:"code"`
	Title  string `json:"title"`
	Description  string    `json:"description"`
	Price     int    `json:"price"`
	Stock int `json:"stock"`
	Status  string `json:"status"`
	Photos string `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
*/
var db []dom.Item

func main() {
	item1 := dom.Item{
		ID:          0,
		Code:        "SAM27324354",
		Title:       "Tablet Samsung Galaxy Tab S7",
		Description: "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       150000,
		Stock:       3,
		Status:      "ACTIVE",
		Photos:      "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	item2 := dom.Item{
		ID:          1,
		Code:        "SAM2555434",
		Title:       "Teclado logitech",
		Description: "Teclado mecanico",
		Price:       20000,
		Stock:       6,
		Status:      "ACTIVE",
		Photos:      "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	useCasesSerce := useCaseService.NewUseCases()

	controllerService := controllerService.NewController(useCasesSerce)
	

	db = append(db, item1, item2)

	r := gin.Default()

	r.GET("/ping", pong)

	r.GET("/api/v1/items", controllerService.getItem)
	r.GET("/api/v1/items/:id", getItemsById)
	r.PUT("/api/v1/items/:id", controllerService.UpdateItem)
	r.POST("/api/v1/items", addItem)
	r.DELETE("/api/v1/items/:id", deleteItem)

	r.Run(port)
}

	func pong(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  "pong",
		})
	}
	
	type responseInfo struct {
		Error bool        `json:"error"`
		Data  interface{} `json:"data"`
	}

	func getItems(c *gin.Context) {
		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  db,
		})
	}

	func saveItem(item *dom.Item) {
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()
		item.ID = obtenerUltimoId(db)+1
		db = append(db, *item)
	}
		
	func addItem(c *gin.Context) {
		request := c.Request
		body := request.Body
	
		var item dom.Item
	
		err := json.NewDecoder(body).Decode(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseInfo{
				Error: true,
				Data:  fmt.Sprintf("invalid json: %s", err.Error()),
			})
			return
		}
		
		if item.Code == "" || item.Title == "" || item.Description == "" || item.Price == 0 || item.Stock == 0 {
			c.JSON(http.StatusBadRequest, responseInfo{
				Error: true,
				Data: "invalid json",
			})
			return
		}

		for _, i:= range db{
			if i.Code == item.Code {
				c.JSON(http.StatusBadRequest, responseInfo{
					Error: true,
					Data:  fmt.Sprintf("invalid json: %s", err.Error()),
				})
				return
			}
		}

		saveItem(&item)
	
		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  item,
		})
	
	}



/*
	func updateItem(ctx *gin.Context) {
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
		err = json.NewDecoder(ctx.Request.Body).Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  err.Error(),
			})
			return
		}
	
		for i, _ := range db {
			if db[i].ID == id {
				db[i].Title = item.Title
				db[i].Description = item.Description
				db[i].Code = item.Code
				db[i].Price = item.Price
	
				ctx.JSON(http.StatusOK, gin.H{"data": db[i]})
				return
			}
		}
	
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontro el item"})
	}
*/


	
	
	func obtenerUltimoId(items []dom.Item) int {
		var ultimoId int
		for _, item := range items {
			if item.ID > ultimoId {
				ultimoId = item.ID
			}
		}
		return ultimoId
	}
	
	func getItemsById(c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  err.Error(),
			})
			return
		}
		for _, i := range db{
			if i.ID == id {
				c.JSON(http.StatusOK, responseInfo{
					Error: false,
					Data: i,
				})
			}
		}

		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "No se encontro el item",
		})
	}

func deleteItem(c *gin.Context){
	
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	for i, v := range db{
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})



}

	
	




