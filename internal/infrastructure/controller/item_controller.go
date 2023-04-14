package controller /* Componentes que no son parte de nuestra app si no que son servicios o herramientas
que utlizamos (estan afuera de nuestra app y nos comunicamos con ellas de alguna forma) */

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	dom "github.com/javmoreno-meli/meli-item-w2/internal/domain"
)

//Para evitar hacer las funciones PUBLICAS podemos crear constructures y hacer que las funciones
//pasen a ser metodos de ese contructor  -> 	Queda pendiente hacer eso

//Nota: que va en Usecases (app)

var Db []dom.Item // esta varaible hara las pases de una base de datos
type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func GetItems(gin *gin.Context) {
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  Db,
	})
}

// Funcion para agregar item
func AddItems(gin *gin.Context) {
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

		//(Note: optimizar el codigo) --> Nueva funcion?
	}
	if item.Code == "" {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "Code is required",
		})
		return
	}
	if item.Title == "" {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "Title is required %s",
		})
		return
	}
	if item.Description == "" {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "Description is required",
		})
		return
	}
	if item.Price == 0 || item.Price < 0 {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "Price is required and need be greater that 0",
		})
		return
	}
	if item.Stock == 0 {
		item.Status = "INACTIVE"
	}
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "Stock must be greater than 0 and must be a number",
		})
		return
	}

	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt

	newId := generateID(Db)
	item.ID = newId
	Db = append(Db, item)
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})

}

// Listar item por id
func GetItemsById(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, value := range Db {
		if value.ID == id {
			gin.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  value,
			})
		}
	}
}

// Actualizar item por ID
func UpdateItems(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := findItemById(id)
	if item == nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("item with ID %d not found", id),
		})
		return
	}

	var updateItem dom.Item
	err = gin.BindJSON(&updateItem)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("error binding json: %s", err.Error()),
		})
		return
	}

	//Reescribir en una funcion (updateFields)
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
	//Por si el stock cambia a 0, entonces se debe poner statos Inactivo
	if updateItem.Stock == 0 {
		item.Status = "INACTIVE"
	}
	//hora de actualizacion
	item.UpdatedAt = time.Now()

	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})

}

// Funcion Delete
func DeleteItem(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := findItemById(id)
	if item == nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("item with ID %d not found", id),
		})
		return
	}

	for i, item := range Db {
		if item.ID == id {
			Db = append(Db[:i], Db[i+1:]...)
			break
		}
	}

	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  "Item delete successfully.",
	})

}

// Funcion para generar ID
// Recibir un SLICE de tipo item
func generateID(items []dom.Item) int {
	maxId := 0
	for i := 0; i < len(items); i++ {
		if items[i].ID > maxId {
			maxId = items[i].ID
		}
	}
	return maxId + 1
}

// Buscar id en slice
func findItemById(id int) *dom.Item {
	for i := range Db {
		if Db[i].ID == id {
			return &Db[i]
		}
	}
	return nil
}
