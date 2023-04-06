package main 

import ( 
	"fmt"
	"net/http" //permite crear u obtener los status //?
	"strconv" 
	"encodigin/json" // json 
	"github.com/gin-gonic/gin" // permite enrutar (metodos get post entro otros)
)

//Puerto en el que correra nuestra API
const port = ":9001"
//Articulos   (las claves del json se obtienen en minusculas como "buena practicas") 
type Item struct {

	Id int				   `json:"id"`
	Code int               `json:"code"`
	Title string           `json:"title"`
	Description string     `json:"description"`
	Price int              `json:"price"`
	Stock int              `json:"stock"`
}

var db []Item // esta varaible hara las pases de una base de datos



func main()  {
	//creemos un item (articulo)
		itemOne := Item{
			Id: 1,
			Code 23,
			Title: "Silla Ergonomica",
			Description: "Silla no solo para sentarse si no para sentarse bien :V"
			Price: 454900,
			Stock: 5,
			 			
		}
		itemTwo := Item{
			Id: 2,
			Code 24,
			Title: "Escritorio de madera",
			Description: "Escritorio para oficina en madera"
			Price: 434900,
			Stock: 10,
			 			
		}
		itemThree := Item{
			Id: 3,
			Code 26,
			Title: "Escritorio de metal",
			Description: "Escritorio para oficina en metal"
			Price: 600000,
			Stock: 4,
			 			
		}

		//ya que estamos, agreguemos los items a nuestra bd (slice)
		db = append (db, itemOne, itemTwo, itemThree)









		

}