package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osalomon89/meli-items-w2/internal/entity"
	"github.com/osalomon89/meli-items-w2/internal/usecase"
)

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) *ItemController {
	return &ItemController{
		itemUsecase: itemUsecase,
	}
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func (ctrl *ItemController) AddItem(c *gin.Context) {
	body := c.Request.Body

	var item entity.Item

	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid json": fmt.Sprint(err.Error())},
		})
		return
	}

	result, err := ctrl.itemUsecase.AddItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, responseInfo{
		Error: false,
		Data:  result})
}

func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	body := c.Request.Body

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	var item entity.Item

	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	result, err := ctrl.itemUsecase.UpdateItemById(item, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Item with id '%d' not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  result})
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := ctrl.itemUsecase.GetItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Item with id '%d' not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item})

}

func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
	}

	item := ctrl.itemUsecase.DeleteItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Item with id '%d' not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  gin.H{fmt.Sprintf("Item with id '%d' deleted", id): item},
	})
}

type Query struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
}

func (ctrl *ItemController) GetItems(c *gin.Context) {
	var q Query

	err := c.ShouldBind(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid query: %s", err.Error()),
		})
		return
	}

	status := q.Status
	limit := q.Limit

	result := ctrl.itemUsecase.GetAllItems(status, limit)

	c.JSON(http.StatusOK, gin.H{
		"totalPages": len(result),
		"data":       result,
	})
}
