package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"meli-items-w2/Internal/domain"
	"meli-items-w2/Internal/usecase"
	"net/http"
	"strconv"
)

type ItemController struct {
	itemUseCase usecase.ItemUsecase
}

func NewItemContrller(itemUseCase usecase.ItemUsecase) *ItemController {
	return &ItemController{}
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

// Función que permite agregar items
func (ctrl *ItemController) AddItem(c *gin.Context) {

	request := c.Request
	var item domain.Item
	err := json.NewDecoder(request.Body).Decode(&item)

	if err := validate.Struct(item); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: false,
			Data:  ctrl.itemUseCase.GetAllItems(),
		})
		return
	}

	item := ctrl.itemUseCase.GeItemByID(id)

	if VerificaRepetido(item) {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido: %S", err.Error()),
		})
		return

	}

	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}

	result, err := ctrl.itemUseCase.AddItem(item)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}

// Funcion para obtener items
func (ctrl *ItemController) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return

	}
	var item domain.Item

	for i, v := range ctrl.db {
		if v.Id == id {
			ctrl.db[i] = item
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "Item no encontrado",
	})

}

// Función que permite acutalizar items
func UpdateItem(c *gin.Context) {
	r := c.Request
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	var item domain.Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err := validate.Struct(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}
	if VerificaRepetido(item) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Código repetido", err.Error()),
		})
		return

	}

	for i, v := range db {
		if v.Id == id {
			db[i] = item
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

// Función que permite obtener items dado un id
func GetItemById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	for _, v := range db {
		if v.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"error": false,
				"data":  v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "Item no encontrado",
	})
}

// Función que permite elimar items
func DeleteItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Sprintf("Json invalido", err.Error()),
		})
		return
	}

	for i, v := range db {
		if v.Id == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "Item no encontrado",
	})
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

// Función que nos dice si el código es repetivo
func VerificaRepetido(item domain.Item) bool {
	for _, i := range db {
		if i.Code == item.Code {
			return true
		}
	}
	return false
}
