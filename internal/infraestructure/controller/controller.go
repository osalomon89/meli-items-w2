package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dom "meli-items-w2/internal/domain"
	"meli-items-w2/internal/usecase"
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

// AddItem Añadir item
func (ctrl *ItemController) AddItem(c *gin.Context) {
	body := c.Request.Body

	var item dom.Item

	err := json.NewDecoder(body).Decode(&item)
	fmt.Sprintf("Entro al decoder")
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid json": fmt.Sprint(err.Error())},
		})
		fmt.Sprintf("Entro al if decoder")
		return
	}

	result := ctrl.itemUsecase.AddItem(item)
	if result == nil {
		fmt.Sprintf("nil de addItem")
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{"invalid param": fmt.Sprint(err.Error())},
		})
		return
	}

	c.JSON(http.StatusCreated, responseInfo{
		Error: false,
		Data:  result})
}

// UpdateItem modificar item
func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	body := c.Request.Body

	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("invalid param: %s", err.Error()): err},
		})
		return
	}

	var item dom.Item

	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("invalid json: %s", err.Error()): err},
		})
		return
	}

	result := ctrl.itemUsecase.UpdateItem(item, id)

	if result == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("Item with id '%d' not found", id): err},
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  result})

}

// GetItemByID Obtener Item por id
func (ctrl *ItemController) GetItemByID(c *gin.Context) {
	// Obtener el ID del parámetro de la URL
	idRequested := c.Param("id")

	// Casteando el param que llega en string a int
	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("invalid param: %s", err.Error()): err},
		})
		return
	}

	item := ctrl.itemUsecase.GetItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("Item with id '%d' not found", id): err},
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item})

}

// DeleteItem Eliminar item
func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	idRequested := c.Param("id")

	id, err := strconv.Atoi(idRequested)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("invalid param: %s", err.Error()): err},
		})
	}

	item := ctrl.itemUsecase.DeleteItem(id)
	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  gin.H{fmt.Sprintf("Item with id '%d' not found", id): err},
		})
		return
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  gin.H{fmt.Sprintf("Item with id '%d' deleted", id): item},
	})
}

// TODO	//	Falta añadir el query param limit
// TODO	//	La respuesta debe seguir la siguiente estructura de campos:
// TODO	//	totalPages: 	El número total de items que contienen resultados para la búsqueda hecha.
// TODO	//	data: 			Un array con los objetos conteniendo los items solicitados en el request.

// ListItem obtener items con filtros
func (ctrl *ItemController) ListItem(c *gin.Context) {
	status := c.Query("status")

	result := ctrl.itemUsecase.ListItem(status)
	if result == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  gin.H{"No data found": result},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
