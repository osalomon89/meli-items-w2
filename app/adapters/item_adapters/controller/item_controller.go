package controller

import (
	"apiRestPractice/app/usecase"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) ItemController {
	return ItemController{
		itemUsecase: itemUsecase,
	}
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func (ctrl ItemController) AddItem(c *gin.Context) {
	req := c.Request
	body := req.Body
	var result map[string]any
	resBody, errRead := io.ReadAll(body)
	if errRead != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error reading body: %s", errRead.Error()),
		})
		return
	}
	err := json.Unmarshal(resBody, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid json: %s", err.Error()),
		})
		return
	}
	errorUseCase := ctrl.itemUsecase.AddItem(result)
	if errorUseCase != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid item, %v", errorUseCase.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  fmt.Sprintf("Item saved!"),
	})
	return
}

func (ctrl ItemController) UpdateItem(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error parse int: %s", errParseInt.Error()),
		})
		return
	}
	req := c.Request
	body := req.Body
	var result map[string]any
	resBody, _ := io.ReadAll(body)
	err := json.Unmarshal(resBody, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid json: %s", err.Error()),
		})
		return
	}

	errorUseCase := ctrl.itemUsecase.UpdateItem(result, id)
	if errorUseCase != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid item, %v", errorUseCase.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  fmt.Sprintf("Item updated!"),
	})
	return
}

func (ctrl ItemController) GetById(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error parse int: %s", errParseInt.Error()),
		})
		return
	}

	errorFound, itemFound := ctrl.itemUsecase.GetById(id)
	itemToString, errMarshal := json.Marshal(itemFound)
	if errMarshal != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error marshal json: %s", errMarshal.Error()),
		})
		return
	}
	if errorFound != nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprint("Id not found"),
		})
		return
	}
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  string(itemToString),
	})
	return
}

func (ctrl ItemController) DeleteById(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error parse int: %s", errParseInt.Error()),
		})
		return
	}

	errorDelete := ctrl.itemUsecase.DeleteById(id)
	if errorDelete != nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error deleting item: %s", errorDelete.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  fmt.Sprint("Item deleted"),
	})
	return
}

func (ctrl ItemController) GetByStatusAndLimit(c *gin.Context) {
	statusReq := c.Query("status")
	limitReq := c.Query("limit")
	limitInt, errAtoi := strconv.Atoi(limitReq)
	if errAtoi != nil {
		c.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error parse int limit: %s", errAtoi.Error()),
		})
		return
	}
	if limitReq == "" {
		limitInt = 0
	}
	itemsRes, errorGet := ctrl.itemUsecase.GetByStatusAndLimit(statusReq, limitInt)
	if errorGet != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error getting response: %s", errorGet.Error()),
		})
		return
	}
	response := make(map[string]any)
	response["totalPages"] = len(itemsRes)
	response["data"] = itemsRes
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  response,
	})
	return
}
