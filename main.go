package main

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/osalomon89/meli-items-w2/internal/controller"
	repo "github.com/osalomon89/meli-items-w2/internal/repository"
	usec "github.com/osalomon89/meli-items-w2/internal/usecase"
)

const port = ":9000"

// var db []dom.Item

func main() {

	itemRepository := repo.NewItemRepository()
	itemUsecase := usec.NewItemUsecase(itemRepository)
	itemController := ctrl.NewItemController(itemUsecase)

	r := gin.Default()

	r.POST("v1/items", itemController.AddItem)
	r.PUT("v1/items/:id", itemController.UpdateItem)
	r.GET("v1/items/:id", itemController.GetItem)
	r.DELETE("v1/items/:id", itemController.DeleteItem) 
	//r.GET("v1/items", getItems) 

	r.Run(port)

}



// Funciones para los endpoints -----------------------------
/*
func addItem(c *gin.Context) {
	request := c.Request
	body := request.Body

	var item dom.Item

	if err := json.NewDecoder(body).Decode(&item); err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	// Empieza logica de negocio
	if codeRepetido(item) {
		
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("El code %s está duplicado", item.Code),
		})
		return
	}

	error_item := initItem(&item)
	if error_item {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  "item nil",
		})
	}
	
	saveItem(item)
	// Termina logica de negocio

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}

func updateItem(c *gin.Context) {
	request := c.Request
	body := request.Body

	idParam := c.Param("id")
	id, err1 := strconv.Atoi(idParam)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err1.Error()),
		})
		return
	}

	var item dom.Item
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}
	
	if codeRepetido(item) {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("El code %s está duplicado", item.Code),
		})
		return
	}

	for pos, val := range db {
		if val.ID == id {
			error_item := actualizarCamposManuales(item, &val)
			if error_item {
				c.JSON(http.StatusBadRequest, responseInfo{
					Error: true,
					Data:  "item nil",
				})
			}
			error_item = actualizarCamposAutomaticos(&val)
			if error_item {
				c.JSON(http.StatusBadRequest, responseInfo{
					Error: true,
					Data:  "item nil",
				})
			}
			db[pos] = val
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  val,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  "Item not found",
	})
}

func getItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: true,
		Data:  "Item not found",
	})

}

func deleteItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	var db_copy []dom.Item
	var encontrado bool

	for _,value := range db {
		if value.ID != id {
			db_copy = append(db_copy, value)
		} else {
			encontrado = true
		}
	}

	if encontrado {
		db = db_copy
		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  db,
		})
		return
	} 

	c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "Item not found",
	})
	
}
*/
/*
func getItems(c *gin.Context){
	status := c.Query("status")
	limitParam := c.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	var db_copy []dom.Item
	var db_copy_sub []dom.Item

	if status == "ACTIVE" {
		for _,value := range db {
			if value.Status == "ACTIVE" {
				db_copy = append(db_copy, value)
			}
		}
		
		db_copy_sub = armarDB(db_copy,limit)
		
		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  db_copy_sub,
		})
	} else if status == "INACTIVE" {
		for _,value := range db {
			if value.Status == "INACTIVE" {
				db_copy = append(db_copy, value)
			}
		}

		db_copy_sub = armarDB(db_copy,limit)

		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  db_copy_sub,
		})

	} else {
		// Si no especifica nada
		c.JSON(http.StatusOK, responseInfo{
			Error: false,
			Data:  db,
		})
	}
}*/







// Funciones auxiliares ---------------------------------------
/*
type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Devuelve true si el code ya existe
func codeRepetido(item dom.Item) bool {
	var repetido bool
	for _, val := range db {
		if val.Code == item.Code {
			repetido = true
		}
	}	
	return repetido
}

// Obtiene el próximo ID libre
func obtenerSiguienteID() int {
	var idSiguiente int
	for _, val := range db {
		if idSiguiente < val.ID {
			idSiguiente = val.ID
		}
	}
	idSiguiente++
	return idSiguiente
}

// Setea el status en funcion del stock
func setStatus(item *dom.Item) bool {
	var error bool
	if item == nil {
		error = true
		return error
	}

	if item.Stock == 0 {
		item.Status = "INACTIVE"
	} else {
		item.Status = "ACTIVE"
	}
	return error
}

func initItem(item *dom.Item) bool{
	var error bool
	if item == nil {
		error = true
		return error
	}
	item.ID = obtenerSiguienteID()
	dt := time.Now()
	item.CreatedAt = dt
	error = actualizarCamposAutomaticos(item)
	return error
}

func actualizarCamposAutomaticos(item *dom.Item) bool{
	var error bool
	if item == nil {
		error = true
		return error
	}
	dt := time.Now()
	item.UpdatedAt = dt
	error = setStatus(item)
	return error
}

func actualizarCamposManuales(item dom.Item, val *dom.Item)bool{
	var error bool
	if val == nil {
		error = true
		return error
	}
	val.Code = item.Code
	val.Title = item.Title
	val.Description = item.Description
	val.Price = item.Price
	val.Stock = item.Stock
	return error
}


func saveItem(item dom.Item){
	db = append(db, item)
}
*/
/*
func armarDB(db_copy []dom.Item, limit int) []dom.Item{
	sort.Slice(db_copy, func(i, j int) bool {
		return db_copy[i].UpdatedAt.After(db_copy[j].UpdatedAt)
	})
	if limit > len(db_copy){
		limit = len(db_copy)
	}
	return db_copy[0:limit]
}
*/