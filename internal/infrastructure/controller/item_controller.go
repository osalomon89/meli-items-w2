package controller /* Componentes que no son parte de nuestra app si no que son servicios o herramientas
que utlizamos (estan afuera de nuestra app y nos comunicamos con ellas de alguna forma) */

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"strconv"
	//"time"

	"github.com/gin-gonic/gin"
	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
	usecase "github.com/javmoreno-meli/meli-item-w2/internal/usecase"
)

//Para evitar hacer las funciones PUBLICAS podemos crear constructures y hacer que las funciones
//pasen a ser metodos de ese contructor  -> 	Queda pendiente hacer eso

type ItemController struct {
	itemUseCase usecase.ItemUseCase
	//db []dom.Item
}

// Constructor
func NewItemController(itemUseCase usecase.ItemUseCase) *ItemController {
	return &ItemController{
		itemUseCase: itemUseCase, //Aca toca colocar (itemUseCase)
	}
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func (c *ItemController) GetItems(gin *gin.Context) {

	items, err := c.itemUseCase.GetItems()
	if err != nil {
		gin.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  items,
	})
}

// Funcion para agregar item
func (c *ItemController) AddItems(gin *gin.Context) {
	//Otra forma : body = gin.Request.Body
	request := gin.Request
	body := request.Body
	var item dom.Item
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido :V %s", err.Error()),
		})
		return
	}

	err = c.itemUseCase.AddItem(&item)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}

// Listar item por id
func (c *ItemController) GetItemsById(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	// llamar al caso de uso para obtener el item
	item, err := c.itemUseCase.GetItemById(id)
	if err != nil {
		gin.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "Item no encontrado",
		})
		return
	}

	// retornar el item encontrado
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}

// Actualizar item por ID
func (c *ItemController) UpdateItems(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid id: %s", gin.Param("id")),
		})
		return
	}

	// Leer el body de la petición y decodificarlo a un item
	request := gin.Request
	body := request.Body
	var item dom.Item
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid JSON: %s", err.Error()),
		})
		return
	}
	// Agregar el id al item
	item.ID = id
	// Actualizar el item en el caso de uso
	var updatedItem *dom.Item
	updatedItem, err = c.itemUseCase.UpdateItem(&item)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}
	// Responder con el item actualizado
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  updatedItem,
	})
}

// Funcion Delete
func (c *ItemController) DeleteItem(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid id: %s", gin.Param("id")),
		})
		return
	}

	// Eliminar el item en el caso de uso
	if err := c.itemUseCase.DeleteItem(id); err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}
	// Responder con mensaje de éxito
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  fmt.Sprintf("Item with id %d has been deleted", id),
	})
}

// Funcion para generar ID
// Recibir un SLICE de tipo item
/* func generateID(items []dom.Item) int {
	maxId := 0
	for i := 0; i < len(items); i++ {
		if items[i].ID > maxId {
			maxId = items[i].ID
		}
	}
	return maxId + 1
} */

//Funcion para verificar que el code no este repetido

// Funcion para cambiar el STATUS segun el STOCK

/* func requeriedFields(item *dom.Item) error {
	if item == nil {
		return errors.New("item is nil")
	}
	if item.Code == "" {
		return errors.New("code is required")
	}
	if item.Title == "" {
		return errors.New("title is required")
	}
	if item.Description == "" {
		return errors.New("description is required")
	}
	if item.Price == 0 || item.Price < 0 {
		return errors.New("price is required and need be greater that 0")

	}
	if item.Stock < 0 {
		return errors.New("stock need be greater that 0")
	}
	return nil
} */

/* func changeItemStatus(item *dom.Item) error {
	if item == nil {
		return errors.New("item is nil")
	}
	if item.Stock == 0 {
		item.Status = "INACTIVE"
		return nil
	}

	item.Status = "ACTIVE"
	return nil
} */

/* func findItemById(id int) *dom.Item {
	for i := range db {
		if db[i].ID == id {
			return &db[i]
		}
	}
	return nil
} */

// codigo unico

/* func verifyCode(code string) bool {

	for i := range db {
		if db[i].Code == code {
			return false
		}
	}
	return true
}
*/

// Actualizar campos (items)  (Comparar original con la que entra (copia))
/* func updateFields(item *dom.Item, updateItem dom.Item) {
	if updateItem.Code != "" {
		item.Code = updateItem.Code
	}
	if updateItem.Title != "" {
		item.Title = updateItem.Title
	}
	if updateItem.Description != "" {
		item.Description = updateItem.Description
	}
	if updateItem.Price != 0 {
		item.Price = updateItem.Price
	}
	if updateItem.Stock != 0 {
		item.Stock = updateItem.Stock
	}
}
*/
