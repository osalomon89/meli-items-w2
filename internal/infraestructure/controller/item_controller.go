package controller

import (
	"encoding/json"
	"fmt"
	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/usecase"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//vienene todas las funciones

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

//msj de error
type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

//nueva funcion
func NewItemController(usecase usecase.ItemUsecase) *ItemController {
	return &ItemController{
		itemUsecase: usecase,
	}
}
/*
// func inicializar para ver su todo funca
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}
*/
/*------metodo getLista inicial de item -------
func (ctrl *ItemController) GetListaInicial(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  ctrl.ItemUsecase.GetListaInicial(),
	})
}
*/


//get all items
func (ctrl *ItemController) GetAllItems(ctx *gin.Context){
	ctx.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data: ctrl.itemUsecase.GetAllItems(),
	})
}



//func additem post
func(ctrl *ItemController) AddItem(ctx *gin.Context){
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

func (ctrl *ItemController) GetItemById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := ctrl.itemUsecase.GetItemById(id)
	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "item not found",
		})
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  item,
	})
}