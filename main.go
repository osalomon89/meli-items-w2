/*
En Mercado Libre trabajamos con articulos de sellers que los venden a traves de nuestro marketplace.
El objetivo de este desafío es realizar una aplicación la cual exponga un API que permita realizar algunas
operaciones de CRUD para cada una de esas dos entidades con algunas reglas de negocio sobre ellas.
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*------declaramos el puerto que escucha -------*/
const port = ":9000"

/*------declaramos la estructura -------*/

type Item struct{
	ID   int    
	Code string
	Title string
	Description string
	Price int
	Stock int
	Status string
	CreatedAt string
	UpdateAt string
}

/*------creamos el slice-------*/
var db []Item

/*------creamos la func-------*/

func main(){
	item1 := Item{
		ID : 1,
		Code : "jddjskdjskdksd",
		Title : "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price : 1234,
		Stock: 54,
		Status: "ACTIVE",
		CreatedAt: "2020-05-10T04:20:33Z",
		UpdateAt: "2020-05-10T05:30:00Z",
	}
	item2 := Item{
		ID : 2,
		Code : "jddjskdjskdksd",
		Title : "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price : 1234,
		Stock: 54,
		Status: "ACTIVE",
		CreatedAt: "2020-05-10T04:20:33Z",
		UpdateAt: "2020-05-10T05:30:00Z",
	}
	item3 := Item{
		ID : 3,
		Code : "jddjskdjskdksd",
		Title : "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price : 1234,
		Stock: 54,
		Status: "ACTIVE",
		CreatedAt: "2020-05-10T04:20:33Z",
		UpdateAt: "2020-05-10T05:30:00Z",
	}

	db = append(db, item1,item2,item3)
	r := gin.Default()

	/*------GETS -------*/

	r.GET("/", index)
	r.GET("v1/items", getItems)
	//r.GET("GET v1/items/{id}", getItemById)
	//r.GET("v1/items?status={status}&limit={limit}", getAllItems)


	/*------POST -------*/

	r.POST("v1/items", addItem)

	/*------PUT-------*/
	r.PUT("v1/items/{id}", updateitem)

	/*------DELETE-------*/
	//r.DELETE("v1/items/{id}", deleteItem)



	r.Run(port)

	/*------MSJ ESCUCHANDO PUERTO -------*/
	log.Println("server listening to the port:", port)


	/*------MSJ DE ERROR -------*/
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}


//w responde: respuesta del servidor al cliente
//r request: peticion del cliente al servidor

//*------estructura de respuesta en error y en data -------*/
type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

//func inicializar para ver su todo funca
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}

/*------metodo GETITEM -------*/
func getItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data": db,
	})
}

//*------1 - metodo POST ADDITEM -------*/
func addItem(ctx *gin.Context) {
	request := ctx.Request

	var item Item
	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data": err.Error(),
		})
		return
	}

	//agregar item
	db = append(db, item)
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data": item,
	})
}

/*------2 - metodo PUT updateitem-------*/
func updateitem(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	var item Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	//aca se usa el id que chiilla
	for i, v := range db {
		if v.ID == id {
			db[i] = item

		}
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data": db,
	})

}

/*------3 - metodo GET ID getid  -------*/
/*------4 - metodo DELETE deleteitem-------*/
/*------5 - metodo GET ALLitems  -------*/