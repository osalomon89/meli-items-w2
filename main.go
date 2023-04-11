/*
En Mercado Libre trabajamos con articulos de sellers que los venden a traves de nuestro marketplace.
El objetivo de este desafío es realizar una aplicación la cual exponga un API que permita realizar algunas
operaciones de CRUD para cada una de esas dos entidades con algunas reglas de negocio sobre ellas.
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*------declaramos el puerto que escucha -------*/
const port = ":9000"

/*------declaramos la estructura -------*/

// se usan las de mayusculas, las minusculas se usan en el postman
// las etiquetas van en el postman
type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}
// *------estructura de respuesta en error y en data -------*/
type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}
/*------creamos el slice-------*/
var articulos []Item

/*------creamos la funcion-------*/
func main() {
	item1 := Item{
		ID:          1,
		Code:        "item1",
		Title:       "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       1234,
		Stock:       54,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	item2 := Item{
		ID:          2,
		Code:        "item2",
		Title:       "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       1234,
		Stock:       54,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	item3 := Item{
		ID:          3,
		Code:        "item3",
		Title:       "Smartphone Samsung s23",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       1234,
		Stock:       54,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	articulos = append(articulos, item1, item2, item3)
	//es el router que crea por default y package de logeo y recover
	r := gin.Default()

	
	/*-----ENDPOINTS-----*/
	//router.GET(el path, n handler))
	/*------GETS -------*/

	r.GET("/", index)
	r.GET("v1/items", getItems)
	r.GET("GET v1/items/:id", getItemById)
	//r.GET("v1/items", getAllItems)

	/*------POST -------*/

	r.POST("v1/items", addItem)

	/*------PUT-------*/
	r.PUT("v1/items/:id", updateItem)

	/*------DELETE-------*/
	r.DELETE("v1/items/:id", deleteItem)

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

// func inicializar para ver su todo funca
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}

/*------metodo getItem -------*/
func getItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  articulos,
	})
}

// func save item para no sobrecargar la funcion additem POST
func saveItem(item *Item) {
	item.CreatedAt = time.Now()
	item.UpdateAt = time.Now()
	//obtenermos
	item.ID = obtenerId()

}

// *------1 - metodo POST ADDITEM -------*/
func addItem(ctx *gin.Context) {
	request := ctx.Request

	var item Item
	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	if codigoRepetido(&item) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Errorf("el code no corresponde, no es unico"),
		})
		return

	}

	//validar status
	item.Status = validateStatus(item.Stock)

	saveItem(&item)

	//que todo este completo obligatoriamente
	if item.Code == "" || item.Title == "" || item.Description == "" || item.Price == 0 || item.Status == "" {
		ctx.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  "invalid json",
		})
		return
	}

	//agregar item
	articulos = append(articulos, item)
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  item,
	})
}

//FUNCIONES -------

// id autoincremental
func obtenerId() int {
	var idSiguiente int
	for _, itemAnterior := range articulos {
		if idSiguiente < itemAnterior.ID {
			idSiguiente = itemAnterior.ID
		}
	}
	//INCREMENTAMOS 1
	idSiguiente += 1
	// incrementar al item
	return idSiguiente
}

// codigo debe ser unico

// funcion repetido
func codigoRepetido(item *Item) bool {
	var repetido bool
	for _, valor := range articulos {
		if valor.Code == item.Code {
			repetido = true
		}
	}
	return repetido
}

// validar status
func validateStatus(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	} else {
		return "INACTIVE"
	}
}

/*------2 - metodo PUT updateitem-------*/
func updateItem(ctx *gin.Context) {
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
	for i, v := range articulos {
		if v.ID == id {
			articulos[i] = item

		}
	}

	item.ID = obtenerId()
	//validar status
	item.Status = validateStatus(item.Stock)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  articulos,
	})

}

/*------3 - metodo GET ID getid  -------*/

func getItemById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range articulos {
		if v.ID == id {
			articulos = append(articulos[:i], articulos[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  articulos,
	})
}

/*------4 - metodo DELETE deleteitem-------*/

func deleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range articulos {
		if v.ID == id {
			articulos = append(articulos[:i], articulos[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  articulos,
	})
}

/*------5 - metodo GET ALLitems  -------*/
