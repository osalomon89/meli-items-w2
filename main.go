package main

import (
	"encoding/json"
	"fmt"
	"net/http" //permite crear u obtener los status (errors) //?
	"strconv"  //Convertir string
	"time"

	"github.com/gin-gonic/gin" // permite enrutar (metodos get post entro otros)
)

// Puerto en el que correra nuestra API
const port = ":9001"

// Articulos   (las claves del json se obtienen en minusculas como "buena practicas")
type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

var db []Item // esta varaible hara las pases de una base de datos

func main() {
	//creemos un item (articulo)
	itemOne := Item{
		ID:          1,
		Code:        "JAV01",
		Title:       "Silla Ergonomica",
		Description: "Silla no solo para sentarse si no para sentarse bien :V",
		Price:       454900,
		Stock:       5,
		Status:      "ACTIVE",
	}
	itemTwo := Item{
		ID:          2,
		Code:        "JAV02",
		Title:       "Escritorio de madera",
		Description: "Escritorio para oficina en madera",
		Price:       434900,
		Stock:       10,
		Status:      "ACTIVE",
	}
	itemThree := Item{
		ID:          3,
		Code:        "JAV03",
		Title:       "Escritorio de metal",
		Description: "Escritorio para oficina en metal",
		Price:       600000,
		Stock:       0,
		Status:      "INACTTIVE",
	}

	//ya que estamos, agreguemos los items a nuestra bd (slice)
	db = append(db, itemOne, itemTwo, itemThree)

	route := gin.Default()
	//Routes
	//listar todos los items en la base de datos (variable db)
	route.GET("/v1/items", getItems)
	//Guardar un item
	route.POST("/v1/items", addItems)
	//Listar item by ID
	route.GET("/v1/items/:id", getItemsById)
	//Actualizar item by ID
	route.PUT("/v1/items/:id", updateItems)
	//Eliminar item by Id
	route.DELETE("/v1/items/:id", deleteItem)

	//Hagamos que nuestras Api corra en el puerto que definimos (9001)
	route.Run(port)
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func getItems(gin *gin.Context) {
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  db,
	})
}

// Funcion para agregar item
func addItems(gin *gin.Context) {
	//Otra forma : body = gin.Request.Body
	request := gin.Request
	body := request.Body
	var item Item
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

	newId := generateID(db)
	item.ID = newId
	db = append(db, item)
	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})

}

// Funcion para generar ID
// Recibir un SLICE de tipo item
func generateID(items []Item) int {
	maxId := 0
	for i := 0; i < len(items); i++ {
		if items[i].ID > maxId {
			maxId = items[i].ID
		}
	}
	return maxId + 1
}

// Listar item por id
func getItemsById(gin *gin.Context) {
	idParam := gin.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, value := range db {
		if value.ID == id {
			gin.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  value,
			})
		}
	}
}

// Actualizar item por ID
func updateItems(gin *gin.Context) {
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

	var updateItem Item
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

// Buscar id en slice
func findItemById(id int) *Item {
	for i := range db {
		if db[i].ID == id {
			return &db[i]
		}
	}
	return nil
}

// Funcion Delete
func deleteItem(gin *gin.Context) {
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

	for i, item := range db {
		if item.ID == id {
			db = append(db[:i], db[i+1:]...)
			break
		}
	}

	gin.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  "Item delete successfully.",
	})

}

//Pendiente
//Agregar funcion que valide si hay otro codigo creado
/*

Refactorizar codigo -->

CREAR FUNCION PARA GUARDAR ITEM (simplifica additem) -
CREAR FUNCION PARA ACTUALIZAR ITEM (simplifica update)


*/
