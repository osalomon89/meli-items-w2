package main

import (
	//"encoding/json" // json
	"encoding/json"
	"fmt"
	"net/http" //permite crear u obtener los status (errors) //?
	"strconv"

	//"strconv"  //Conviertir string
	"time"

	"github.com/gin-gonic/gin" // permite enrutar (metodos get post entro otros)
)

// Puerto en el que correra nuestra API
const port = ":9001"

// Articulos   (las claves del json se obtienen en minusculas como "buena practicas")
type Item struct {
	ID          int    `json:"id"` //El id no se debe mandar aca si no automaticamente
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	//Listar Items by ID
	route.GET("/v1/items/:id", getItemsById)
	//Actualizar Items by ID
	route.PUT("/v1/items/:id", updateItems)

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

		//(Note: optimizar el codigo)
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

// Actualizar item
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

//Pendiente
//Agregar funcion que valide si hay otro codigo creado
//Crear funcion para agregar el status y la fecha de creacion o actualizacion
