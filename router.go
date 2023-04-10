package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

// TODO: make Items code unique

func main() {
	router := gin.Default()

	// Define "DB"
	items := make(map[int]Item)

	// Define API routes here

	router.GET("v1/items", func(c *gin.Context) {
		status := c.Query("status")                  // get status query parameter
		limit, err := strconv.Atoi(c.Query("limit")) // get limit query parameter
		if err != nil {
			limit = len(items) // default limit is all items
		}
		var filteredItems []Item     // create an empty slice for the filtered items
		for _, item := range items { // loop through all the items
			if item.Status == status { // if the item's status matches the query parameter
				filteredItems = append(filteredItems, item) // add the item to the filtered list
			}
		}
		// If there are no items
		if len(items) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No items found"})
			return
		}
		// If there are no items with the matching status
		if len(filteredItems) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No items found with the given status"})
			return
		}
		// If the limit is greater than the number of filtered items
		if limit > len(filteredItems) {
			limit = len(filteredItems)
		}
		c.JSON(http.StatusOK, gin.H{
			"items": filteredItems[:limit], // return only the first `limit` number of filtered items
		})
	})

	// GET: returns the item with the matching id or err
	router.GET("v1/items/:id", func(c *gin.Context) {
		// Extract the ID parameter from the URL path
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		fmt.Println(items[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Look up the item in the map
		item, ok := items[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		// Return the item
		c.JSON(http.StatusOK, item)
	})

	// POST: creates item and returns it or errs
	router.POST("v1/items", func(c *gin.Context) {
		var item Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Assign a unique ID to the item
		item.ID = len(items) + 1

		item.Created = time.Now()
		// Add the item to the map
		items[item.ID] = item

		// Return the created item
		c.JSON(http.StatusCreated, item)
	})

	// PUT: updates the editable attributes of item with ID or errs
	router.PUT("v1/items/:id", func(c *gin.Context) {
		// Extract the ID parameter from the URL path
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Look up the item in the map
		item, ok := items[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		// Update the item with the new values
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item.Updated = time.Now()
		// Store the updated item back in the map
		items[id] = item

		// Return the updated item
		c.JSON(http.StatusOK, item)
	})

	// DELETE: removes the item of ID from the map ("DATABASE")
	router.DELETE("v1/items/:id", func(c *gin.Context) {
		// Extract the ID parameter from the URL path
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Look up the item in the map
		_, ok := items[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		// Delete the item from the map
		delete(items, id)

		// Return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
	})

	// Run the server
	router.Run(":8080")
}
