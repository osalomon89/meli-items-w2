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
	"sort"
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
	Data  interface{} `json:"data"`
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
	r.GET("v1/listaInicial", getListaInicial)
	r.GET("v1/items/:id", getItemById)
	r.GET("v1/items", getAllItems)

	/*------POST -------*/
	r.POST("v1/items", addItem)

	/*------PUT-------*/
	r.PUT("v1/items/:id", updateItem)

	/*------DELETE-------*/
	r.DELETE("v1/items/:id", deleteItem)

	/*------MSJ ESCUCHANDO PUERTO -------*/
	r.Run(port)
	log.Println("server listening to the port:", port)

	/*------MSJ DE ERROR -------*/
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

// func inicializar para ver su todo funca
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}

/*------metodo getItem -------*/
func getListaInicial(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  articulos,
	})
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

	//validacion de informacion completa
	//que todo este completo obligatoriamente (se la robe a santi)
	err = informacionCompleta(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  "invalid json",
		})
	}

	//ver codigo repetido
	if codigoRepetido(&item) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Errorf("el code no corresponde, no es unico"),
		})
		return

	}

	//lamamos las funciones
	actualizarCreateAt(&item)
	actualizarUpdateAt(&item)
	validateStatus(&item)
	actualizarId(&item)

	appendItemToArticulos(item)

	//agregar item
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  item,
	})
}

/*
/*------2 - metodo PUT updateitem-------
*/
func updateItem(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	_, err := strconv.Atoi(idParam)
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
	//CHECKEO DE CODE UNICO
	if codigoRepetido(&item) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Errorf("el code no corresponde, no es unico"),
		})
		return

	}
	//lamamos funciones
	updateItemNuevo(item)

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
		ctx.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, v := range articulos {
		if v.ID == id {
			ctx.JSON(http.StatusOK, ResponseInfo{
				Error: false,
				Data:  v,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, ResponseInfo{
		Error: true,
		Data:  "item not found",
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
//meteria los slice los active, los ordenaria y filtraria por limit

func getAllItems(ctx *gin.Context) {
	status := ctx.Query("status")
	limitParam := ctx.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid Param: %s", err.Error()),
		})
		return
	}

	var articulos_copy []Item
	var articulos_copy_sub []Item

	if status == "ACTIVE" {
		for _, value := range articulos {
			if value.Status == "ACTIVE" {
				articulos_copy = append(articulos_copy, value)
			}
		}
		sort.Slice(articulos_copy, func(i, j int) bool {
			return articulos_copy[i].UpdateAt.After(articulos_copy[j].UpdateAt)
		})
		if limit > len(articulos_copy) {
			limit = len(articulos_copy)
		}
		articulos_copy_sub = articulos_copy[0:limit]

		ctx.JSON(http.StatusOK, ResponseInfo{
			Error: false,
			Data: articulos_copy_sub,
		})
	} else if status == "INACTIVE" {
		for _, value := range articulos {
			if value.Status == "INACTIVE" {
				articulos_copy = append(articulos_copy, value)
			}
		}
		sort.Slice(articulos_copy, func(i, j int) bool {
			return articulos_copy[i].UpdateAt.After(articulos_copy[j].UpdateAt)
		})
		if limit > len(articulos_copy) {
			limit = len(articulos_copy)
		}
		articulos_copy_sub = articulos_copy[0:limit]
		ctx.JSON(http.StatusOK, ResponseInfo{
			Error: false,
			Data:  articulos_copy_sub,
		})
	} else {
		// Si no especifica nada
		ctx.JSON(http.StatusOK, ResponseInfo{
			Error: false,
			Data:  articulos,
		})
	}
}
