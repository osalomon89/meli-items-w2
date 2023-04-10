package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	id          int
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	status      string
	created_at  string
	updated_at  string
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

const port = ":8080"

var articulos []Item

var countId int = 0

func main() {

	i1 := Item{
		Code:        "item1",
		Title:       "Balon de futbol",
		Description: "Balon de futbol para jugar en campo sintetico",
		Price:       20,
		Stock:       15,
	}

	i2 := Item{
		Code:        "item2",
		Title:       "Balon de volleyball",
		Description: "Balon de volleyball para jugar en la arena",
		Price:       10,
		Stock:       30,
	}

	i3 := Item{
		Code:        "item3",
		Title:       "Pelota de golf",
		Description: "Pelota de golf de color blanco",
		Price:       5,
		Stock:       0,
	}

	saveItem(&i1)
	saveItem(&i2)
	saveItem(&i3)

	server := gin.Default()

	server.POST("v1/items", addItem)
	server.PUT("v1/items/:id", updateItem)
	server.GET("v1/items/:id", getItem)
	server.DELETE("v1/items/:id", deleteItem)
	server.GET("v1/items", getItems)

	server.Run(port)

}

func addItem(c *gin.Context) {
	body := c.Request.Body

	var item Item

	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid json": fmt.Sprint(err.Error())},
		})
		return
	}

	err = validateCode(item.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid param": fmt.Sprint(err.Error())},
		})
		return
	}

	saveItem(&item)

	c.JSON(http.StatusCreated, responseInfo{
		Error: false,
		Data: gin.H{
			"id":          item.id,
			"code":        item.Code,
			"title":       item.Title,
			"description": item.Description,
			"price":       item.Price,
			"stock":       item.Stock,
			"status":      item.status,
			"created_at":  item.created_at,
			"updated_at":  item.updated_at,
		}})
}

func saveItem(i *Item) {
	countId++
	i.id = countId

	i.status = validateStatus(i.Stock)

	t := time.Now().Format(time.RFC3339)
	i.created_at = fmt.Sprint(t)
	i.updated_at = fmt.Sprint(t)

	articulos = append(articulos, *i)
}

func validateCode(code string) error {
	key := 0
	for key < len(articulos) {
		if articulos[key].Code == code {
			return errors.New(fmt.Sprintf("The code '%s' already exists", code))
		}
		key++
	}
	return nil
}

func validateStatus(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}

func updateItem(c *gin.Context) {
	body := c.Request.Body

	var item Item
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid json": fmt.Sprint(err.Error())},
		})
		return
	}

	handlingInvalidParam := func(e error) {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"Invalid param": fmt.Sprintf(e.Error())},
		})
	}

	idRequested := c.Param("id")
	id, err := strconv.Atoi(idRequested)
	if err != nil {
		handlingInvalidParam(err)
		return
	}

	key := 0
	for key < len(articulos) {
		if articulos[key].id == id {
			err = validateCode(item.Code)
			if err != nil && articulos[key].Code != item.Code {
				handlingInvalidParam(err)
				return
			}

			articulos[key].Code = item.Code
			articulos[key].Title = item.Title
			articulos[key].Description = item.Description
			articulos[key].Price = item.Price
			articulos[key].Stock = item.Stock
			articulos[key].status = validateStatus(item.Stock)
			articulos[key].updated_at = fmt.Sprintf(time.Now().Format(time.RFC3339))

			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data: gin.H{
					"id":          articulos[key].id,
					"code":        articulos[key].Code,
					"title":       articulos[key].Title,
					"description": articulos[key].Description,
					"price":       articulos[key].Price,
					"stock":       articulos[key].Stock,
					"status":      articulos[key].status,
					"created_at":  articulos[key].created_at,
					"updated_at":  articulos[key].updated_at,
				}})
			return
		}
		key++
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  fmt.Sprintf("Item with id '%d' not found", id),
	})

}

func getItem(c *gin.Context) {

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"Invalid param": fmt.Sprintf(err.Error())},
		})
		return
	}

	key := 0
	for key < len(articulos) {
		if articulos[key].id == id {
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data: gin.H{
					"id":          articulos[key].id,
					"code":        articulos[key].Code,
					"title":       articulos[key].Title,
					"description": articulos[key].Description,
					"price":       articulos[key].Price,
					"stock":       articulos[key].Stock,
					"status":      articulos[key].status,
					"created_at":  articulos[key].created_at,
					"updated_at":  articulos[key].updated_at,
				}})
			return
		}
		key++
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  fmt.Sprintf("Item with id '%d' not found", id),
	})

}

func deleteItem(c *gin.Context) {
	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"Invalid param": fmt.Sprintf(err.Error())},
		})
	}

	key := 0
	for key < len(articulos) {
		if articulos[key].id == id {
			item := articulos[key]
			articulos = append(articulos[:key], articulos[key+1:]...)
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  gin.H{fmt.Sprintf("Item with id '%d' deleted", id): item},
			})
			return
		}
		key++
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  fmt.Sprintf("Item with id '%d' not found", id),
	})
}

type Query struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
}

func getItems(c *gin.Context) {

	var q Query

	err := c.ShouldBind(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"Invalid query": fmt.Sprintf(err.Error())},
		})
		return
	}

	status := q.Status
	limit := q.Limit

	if limit <= 0 {
		limit = 10
	} else if limit > 20 {
		limit = 20
	} else if limit > len(articulos) {
		limit = len(articulos)
	}

	var itemsToshow []Item
	if len(status) != 0 {
		for k, v := range articulos {
			if v.status == status {
				itemsToshow = append(itemsToshow, v)
			}
			if k == limit-1 {
				break
			}
		}
	} else {
		itemsToshow = append(articulos[:limit])
	}

	c.JSON(http.StatusOK, gin.H{
		"totalPages": len(itemsToshow),
		"data":       itemsToshow,
	})

}
