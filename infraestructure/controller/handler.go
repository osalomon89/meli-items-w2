package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dom "meli-items-w2/domain"
	"net/http"
	"strconv"
)

type controller struct {
}

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func NewController() *controller {
	return &controller{}
}

// AddItem Añadir item
func (ctrl *controller) AddItem(c *gin.Context) {
	var newSliceItem []dom.Item

	// Manejando el error
	if err := c.BindJSON(&newSliceItem); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	saveItem(newSliceItem)
	c.IndentedJSON(http.StatusCreated, newSliceItem)

}

// UpdateItem modificar item
func (ctrl *controller) UpdateItem(c *gin.Context) {
	//var itemFound Item
	var itemToUpdate dom.Item
	id := c.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if error := c.BindJSON(&itemToUpdate); error != nil {
		c.JSON(http.StatusNotFound, error)
		return
	}
	/*
		var itemToUpdatePtr *dom.Item

		for i := range itemsDB {
			if itemsDB[i].Id == idInt {
				itemToUpdatePtr = &itemsDB[i]
				break
			}
		}

		if itemToUpdatePtr == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
			return
		}

		if itemToUpdate.Stock > 0 {
			itemToUpdate.Status = "ACTIVE"
		} else {
			itemToUpdate.Status = "INACTIVE"
		}

		itemToUpdatePtr.Code = itemToUpdate.Code
		itemToUpdatePtr.Title = itemToUpdate.Title
		itemToUpdatePtr.Description = itemToUpdate.Description
		itemToUpdatePtr.Price = itemToUpdate.Price
		itemToUpdatePtr.Stock = itemToUpdate.Stock
		itemToUpdatePtr.Status = itemToUpdate.Status
		itemToUpdatePtr.UpdatedAt = time.Now()

		c.JSON(http.StatusOK, itemToUpdatePtr)
	*/

	c.JSON(http.StatusOK, "todo bien")

}

// GetItemByID Obtener Item por id
func (ctrl *controller) GetItemByID(c *gin.Context) {
	// Obtener el ID del parámetro de la URL
	id := c.Param("id")

	// Para guardar el item si se encuentra
	var itemFound dom.Item

	// Casteando el param que llega en string a int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	// Buscamos el id que necesitamos
	for _, item := range itemsDB {
		if item.Id == idInt {
			itemFound = item
			break
		}
	}

	// Retornamos ok si encontramos el id, no es necesario igular a true en la condición ya que en Go porque el valor booleano es en sí una condición en el caso if found {...}
	// Se puede evitar el uso de la variable bandera si directamente preguntamos por el valor por default de la struct Item
	if itemFound != (dom.Item{}) {
		c.JSON(http.StatusOK, itemFound)
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
	}

}

// DeleteItem Eliminar item
func (ctrl *controller) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, "no se encuentra el id solicitado")
}

// TODO	//	La respuesta debe seguir la siguiente estructura de campos:
// TODO	//	totalPages: 	El número total de items que contienen resultados para la búsqueda hecha.
// TODO	//	data: 			Un array con los objetos conteniendo los items solicitados en el request.

// ListItem obtener items con filtros
func (ctrl *controller) ListItem(c *gin.Context) {
	status := c.Query("status")

	var dbFiltered []dom.Item

	if status != "ACTIVE" && status != "INACTIVE" && status != "ALL" {
		c.Error(fmt.Errorf("status inválido"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, item := range itemsDB {
		if item.Status == status {
			dbFiltered = append(dbFiltered, item)
		}

	}

	if len(itemsDB) == 0 {
		c.Error(fmt.Errorf("no hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if len(dbFiltered) == 0 {
		c.Error(fmt.Errorf("no hay items disponibles"))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, dbFiltered)

}
