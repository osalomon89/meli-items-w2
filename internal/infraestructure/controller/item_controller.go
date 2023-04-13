package controller

import (
	"encoding/json"
	"fmt"
	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/domain/port"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//vienene todas las funciones
//esta estructura recibe una interface de usecase para poder comunicarse

type ItemController struct {
	itemUsecase port.ItemUsecase
}

// msj de error
type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// constructor --- unica funcion
func NewItemController(usecase port.ItemUsecase) *ItemController {
	return &ItemController{
		itemUsecase: usecase,
	}
}


/*------------GET-------------*/


// func inicializar para ver su todo funca
func (ctrl *ItemController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}

/*------metodo getLista inicial de item -------*/

func (ctrl *ItemController) GetListaInicial(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  ctrl.itemUsecase.GetAllItems(),
	})
}

// get all items
func (ctrl *ItemController) GetAllItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  ctrl.itemUsecase.GetAllItems(),
	})
}


func (ctrl *ItemController) GetItemById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := ctrl.itemUsecase.GetItemById(id)
	if item == nil {
		ctx.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "item not found",
		})
	}

	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}
/*----------POST------------*/
// func additem post
func (ctrl *ItemController) AddItem(ctx *gin.Context) {
	request := ctx.Request

	var item domain.Item
	err := json.NewDecoder(request.Body).Decode(&item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	result, err := ctrl.itemUsecase.AddItem(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("error saving item: %s", err.Error()),
		})
		return
	}
	//agregar item
	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  result,
	})
}

/*----------------PUT----------------*/

// ctrl con puntero a itemcontroler
func (ctrl *ItemController) UpdateItem(ctx *gin.Context) {
	// aca dentro solo corresponde algo, no lo logico del negocio
	r := ctx.Request
	idParam := ctx.Param("id")
	

	id, err := strconv.Atoi(idParam)
	if err != nil {
		
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Errorf("invalid ID parameter: %s", err.Error()),
		})
		return
	}
	var item domain.Item

	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println("VER decodifica body")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  fmt.Errorf("error decoding request body: %s", err.Error()),
		})
		return
	}
	item.ID = id

	itemActualizado,err := ctrl.itemUsecase.UpdateItem(item)
	if err!=nil{
		
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  itemActualizado,
	})
}