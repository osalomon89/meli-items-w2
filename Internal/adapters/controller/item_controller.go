package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"meli-items-w2/Internal/entity"
	"meli-items-w2/Internal/usecase"
	"net/http"
	"strconv"
)

// Mensaje bienvenida
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
type ResponseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type ItemController struct {
	itemUseCase usecase.ItemUsecase
}

func NewItemContrller(itemUseCase usecase.ItemUsecase) *ItemController {
	return &ItemController{
		itemUseCase: itemUseCase,
	}
}

// Función que permite agregar items
func (ctrl *ItemController) AddItem(c *gin.Context) {
	request := c.Request

	var item entity.Item

	if err := json.NewDecoder(request.Body).Decode(&item); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: false,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	result, err := ctrl.itemUseCase.AddItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido: %S", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  result,
	})
}

// Función que permite acutalizar items
func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	r := c.Request
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido: %S", err.Error()),
		})
		return
	}
	var item entity.Item

	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}
	result, err := ctrl.itemUseCase.UpdateItemById(item, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}
	if err == nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Error al actualizar el elemento: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  result})
}

// Funcion para obtener items
func (ctrl *ItemController) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Item no enconttrado", err.Error()),
		})
		return
	}

	item := ctrl.itemUseCase.GetItemById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Item no encpntrado", err.Error()),
		})
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  fmt.Sprintf("El item es", item),
	})
}

// Función que permite obtener items dado un id
func (ctrl *ItemController) GetItemById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Id invalido", err.Error()),
		})
		return
	}
	item := ctrl.itemUseCase.GetItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error: true,
			Data:  ("No se encontró ningún elemento con el ID %d", id)
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  fmt.Sprintf("El id del item es", item),
	})
}

// Función que permite elimar items
func (crtl *ItemController) DeleteItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	item := crtl.itemUseCase.DeleteItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("id incorrecto", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Data:  fmt.Sprintf("El item fue eliminado correctamente", item),
	})
}
