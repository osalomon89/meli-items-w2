// resolver problema del go
// go mod init 'yourmodulename'
// go get 'yourpckagename'
package main

import (
	"fmt"
	"gigigarino/challengeMELI/internal/infraestructure/controller"
	"gigigarino/challengeMELI/internal/infraestructure/repository/in-memory"
	"gigigarino/challengeMELI/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//llamar las funciones del controller

//crear estructura

//func newbookcontroler () bookcontroller {
//todos metodos vamos a ir pasandole los metodos de las funciones
//convierten la funciones publicas a metodos publicos
//}

/*------declaramos el puerto que escucha -------*/
const port = ":9000"

func main() {

	//PUEDO PONER LOS ITEMS
	/*
	item1 := Item{
		ID:          0,
		Code:        "SAM27324000",
		Title:       "Smartphone Samsung s23 Ultra",
		Description: "Smarthphone Galaxy s23 S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
		Price:       400000,
		Stock:       54,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	item2 := Item{
		ID:          1,
		Code:        "SAM27324001",
		Title:       "Smartphone Samsung a10",
		Description: "Smarthphone Galaxy a10S sin Pen 10 pulgadas y 2GB de memoria RAM",
		Price:       120000,
		Stock:       120,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	item3 := Item{
		ID:          2,
		Code:        "SAM27324002",
		Title:       "Microondas Samsung AB456",
		Description: "Microondas Samsung AB456 con cooking iron and roasting function",
		Price:       156000,
		Stock:       20,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	*/

	//mysql
	/*
	conn, err := mysql.RepositoryConn()
	if err != nil {
		log.Fatal(err)
	}
	*/
	// se crea la instancia de gin engine, 
	//crea un entitador  con los middleware de logger y recovery por defecto
	r := gin.Default()

	//va el repo/usecase/controller
	repo := repository.NewItemRepository()
	usecase := usecase.NewItemUsecase(repo)
	ctrl := controller.NewItemController(usecase)

	fmt.Println(ctrl)

	

	//van los endpoint, se crean las rutas
	// esto iria en la carpeta web
	/*------GETS -------*/
	r.GET("/", ctrl.Index)
	r.GET("v1/listaInicial", ctrl.GetListaInicial)
	r.GET("v1/items/:id", ctrl.GetItemById)
	r.GET("v1/items", ctrl.GetAllItems)

	/*------POST -------*/
	r.POST("v1/items", ctrl.AddItem)

	/*------PUT-------*/
	r.PUT("v1/items/:id", ctrl.UpdateItem)

	/*------DELETE-------*/
	r.DELETE("v1/items/:id", ctrl.DeleteItem)


	/*------MSJ ESCUCHANDO PUERTO -------*/
	r.Run(port)
	log.Println("server listening to the port:", port)

	/*------se crea el servidor con el metodo run de gin -------*/
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
	
}