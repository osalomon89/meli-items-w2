package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/osalomon89/meli-items-w2/internal/usecase"
	dom "github.com/osalomon89/meli-items-w2/internal/domain"
)

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) *ItemController {
	return &ItemController{
		itemUsecase: itemUsecase,
	}
}

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// --------------------------------------------------------------------------

func (ctrl *ItemController) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := ctrl.itemUsecase.GetItemByID(id)

	if item == nil {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "Item not found",
		})
	} 
	
	c.JSON(http.StatusNotFound, responseInfo{
		Error: false,
		Data:  item,
	})

}

func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item := ctrl.itemUsecase.DeleteItemByID(id)

	/* Logica de negocio
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
	*/
	if !item {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  "Item not found",
		})
	}

	c.JSON(http.StatusNotFound, responseInfo{
		Error: false,
		Data:  ctrl.itemUsecase.GetAllItems(),
})
}

func (ctrl *ItemController) AddItem(c *gin.Context) {
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

	item1,my_err := ctrl.itemUsecase.AddItemByItem(&item)

	if my_err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Errorf("HUBO UN ERROR: %s",my_err.Error()),
		})
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  &item1,
	})
}

func (ctrl *ItemController) UpdateItem(c *gin.Context) {
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
	

	item1,my_err := ctrl.itemUsecase.UpdateItemByItem(id,&item)

	if my_err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Errorf("HUBO UN ERROR: %s",my_err.Error()),
		})
	}

	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  &item1,
	})
}