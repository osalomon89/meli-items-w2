package main

import (
	"encoding/json"
	"net/http"

	"github.com/NitroML/meli-items-w2/domain"
	"github.com/NitroML/meli-items-w2/infra/create"
	"github.com/NitroML/meli-items-w2/infra/getall"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define the v1 group
	v1 := r.Group("/v1")

	// Define the /items endpoint and register the ItemsGetAll controller function
	v1.GET("/items", func(c *gin.Context) {
		items, err := getall.ItemsGetAll()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, items)
	})

	// Define the /items endpoint for creating new items
	v1.POST("/items", func(c *gin.Context) {
		var item domain.Item
		err := json.NewDecoder(c.Request.Body).Decode(&item)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		newItem, err := create.ItemCreate(item)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, newItem)
	})

	// Start the server
	r.Run(":8080")
}
