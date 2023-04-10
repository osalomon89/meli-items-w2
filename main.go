package main

import (
	//"encoding/json" // json
	"encoding/json"
	"fmt"
	"net/http" //permite crear u obtener los status (errors) //?

	"github.com/gin-gonic/gin" // permite enrutar (metodos get post entro otros)
	//"strconv"                  //Conviertir string
)

// Puerto en el que correra nuestra API
const port = ":9001"

// Articulos   (las claves del json se obtienen en minusculas como "buena practicas")
type Item struct {
	Id          int    `json:"id"` //El id no se debe mandar aca si no automaticamente
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Status      string
	//CreatedAt  time.Time
	//UpdatedAt  time.Time
}

var db []Item // esta varaible hara las pases de una base de datos

func main() {
	//creemos un item (articulo)
	itemOne := Item{
		Id:          1,
		Code:        "JAV01",
		Title:       "Silla Ergonomica",
		Description: "Silla no solo para sentarse si no para sentarse bien :V",
		Price:       454900,
		Stock:       5,
		Status:      "ACTIVE",
	}
	itemTwo := Item{
		Id:          2,
		Code:        "JAV02",
		Title:       "Escritorio de madera",
		Description: "Escritorio para oficina en madera",
		Price:       434900,
		Stock:       10,
	}
	itemThree := Item{
		Id:          3,
		Code:        "JAV03",
		Title:       "Escritorio de metal",
		Description: "Escritorio para oficina en metal",
		Price:       600000,
		Stock:       4,
	}

	//ya que estamos, agreguemos los items a nuestra bd (slice)
	db = append(db, itemOne, itemTwo, itemThree)

	route := gin.Default()
	//Routes
	//listar todos los items en la base de datos (variable db)
	route.GET("/api/v1/items", getItems)
	//Guardar un item

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

func addItems(gin *gin.Context) {
	//Otra forma : body = gin.Request.Body
	request := gin.Request
	body := request.Body

	var item Item
	error := json.NewDecoder(body).Decode(&item)
	if error != nil {
		gin.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Json invalido :V %s", error.Error()),
		})
	}
}

//Preguntas?

//Como hacer un Id unico sin una base de datos(:V For con condiciones?) y que se generen automaticamente (ramdom?) -> Crear un contador

//
