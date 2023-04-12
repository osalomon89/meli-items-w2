package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
		Data:  item})
}

func saveItem(i *Item) {
	countId++
	i.Id = countId

	i.Status = validateStatus(i.Stock)

	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()

	articulos = append(articulos, *i)
}

func validateCode(code string) error {

	for _, v := range articulos {
		if v.Code == code {
			return errors.New(fmt.Sprintf("The code '%s' already exists", code))
		}
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

	var item Item
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid json": fmt.Sprint(err.Error())},
		})
		return
	}

	for _, v := range articulos {
		if v.Id == id {
			err = validateCode(item.Code)
			if err != nil && v.Code != item.Code {
				handlingInvalidParam(err)
				return
			}

			v.Code = item.Code
			v.Title = item.Title
			v.Description = item.Description
			v.Price = item.Price
			v.Stock = item.Stock
			v.Status = validateStatus(item.Stock)
			v.UpdatedAt = time.Now()

			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  v})
			return
		}
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

	for _, v := range articulos {
		if v.Id == id {
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  v})
			return
		}
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

	for k, v := range articulos {
		if v.Id == id {
			articulos = append(articulos[:k], articulos[k+1:]...)
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  gin.H{fmt.Sprintf("Item with id '%d' deleted", id): v},
			})
			return
		}
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

	sort.SliceStable(articulos, func(i, j int) bool {
		return articulos[i].UpdatedAt.After(articulos[j].UpdatedAt)
	})
	var itemsToshow []Item

	if len(status) != 0 {
		for k, v := range articulos {
			if v.Status == status {
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
