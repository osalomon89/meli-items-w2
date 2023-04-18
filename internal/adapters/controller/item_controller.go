package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osalomon89/meli-items-w2/internal/adapters/controller/presenter"
	"github.com/osalomon89/meli-items-w2/internal/entity"
	"github.com/osalomon89/meli-items-w2/internal/usecase"
)

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) ItemController {
	return ItemController{
		itemUsecase: itemUsecase,
	}
}

type itemRequestJSON struct {
	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock"`
}

func (ic *ItemController) AddItem(c *gin.Context) {

	var itemRequest itemRequestJSON

	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid json: %s", err.Error()),
		})
		return
	}

	item := entity.Item{
		Code:        itemRequest.Code,
		Title:       itemRequest.Title,
		Description: itemRequest.Description,
		Price:       itemRequest.Price,
		Stock:       itemRequest.Stock,
	}

	result, err := ic.itemUsecase.AddItem(item)
	if err != nil {
		var msg string
		var status int

		duplicatedError := new(entity.ItemAlreadyExist)
		ok := errors.As(err, duplicatedError)
		if ok {
			status = http.StatusBadRequest
			msg = duplicatedError.Error()
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
		}

		c.JSON(status, presenter.ApiError{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusCreated, presenter.ItemResponse{
		Error: false,
		Data:  presenter.Item(result)})
}

func (ic *ItemController) UpdateItem(c *gin.Context) {

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	var itemRequest itemRequestJSON

	if err = c.ShouldBindJSON(&itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	item := entity.Item{
		Code:        itemRequest.Code,
		Title:       itemRequest.Title,
		Description: itemRequest.Description,
		Price:       itemRequest.Price,
		Stock:       itemRequest.Stock,
	}

	result, err := ic.itemUsecase.UpdateItemById(item, id)
	if err != nil {
		var status int
		var msg string

		duplicated := new(entity.ItemAlreadyExist)
		notExists := new(entity.ItemNotFound)

		if errors.As(err, duplicated) {
			status = http.StatusBadRequest
			msg = duplicated.Message
		} else if errors.As(err, notExists) {
			status = http.StatusNotFound
			msg = notExists.Message
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
		}
		c.JSON(status, presenter.ApiError{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemResponse{
		Error: false,
		Data:  presenter.Item(result)})
}

func (ic *ItemController) GetItem(c *gin.Context) {

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item, err := ic.itemUsecase.GetItemById(id)
	if err != nil {
		var status int
		var msg string

		notExists := new(entity.ItemNotFound)

		if errors.As(err, notExists) {
			status = http.StatusNotFound
			msg = notExists.Message
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
		}
		c.JSON(status, presenter.ApiError{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemResponse{
		Error: false,
		Data:  presenter.Item(item)})

}

func (ic *ItemController) DeleteItem(c *gin.Context) {

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item, err := ic.itemUsecase.DeleteItemById(id)
	if err != nil {
		var status int
		var msg string

		notExists := new(entity.ItemNotFound)

		if errors.As(err, notExists) {
			status = http.StatusNotFound
			msg = notExists.Message
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
		}
		c.JSON(status, presenter.ApiError{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemDeletedResponse{
		Error:   false,
		Message: fmt.Sprintf("Item with id '%d' deleted", id),
		Data:    presenter.Item(item),
	})
}

func (ic *ItemController) GetItems(c *gin.Context) {

	status := c.Query("status")
	lim := c.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(lim)

	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	result, err := ic.itemUsecase.GetAllItems(status, limit)
	if err != nil {
		var status int
		var msg string

		notExists := new(entity.ItemNotFound)

		if errors.As(err, notExists) {
			status = http.StatusNotFound
			msg = notExists.Message
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
		}
		c.JSON(status, presenter.ApiError{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemsResponse{
		Error: false,
		Data:  presenter.Items(result),
	})
}
