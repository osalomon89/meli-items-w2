package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dom "github.com/osalomon89/neocamp-meli-w2/domain"

	domainPorts "github.com/osalomon89/neocamp-meli-w2/domain/ports"
)

type controller struct {
	ucService domainPorts.UseCasesService
	ucRepository domainPorts.Repository
}

func NewController(ucService domainPorts.UseCasesService ) *controller{
	controllerService := controller{
		ucService: ucService,
	}
	return &controllerService
}


func (ctrl *controller) UpdateItem(ctx *gin.Context) {
		
		idParam := ctx.Param("id")
		 idInt, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  err.Error(),
			})
			return
		}
	
		var item dom.Item
		err = json.NewDecoder(ctx.Request.Body).Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  err.Error(),
			})
			return
		}

		item.ID = idInt
		
		err = ctrl.ucService.UpdateItem(item)

		

	/*
		for i, _ := range db {
			if db[i].ID == id {
				db[i].Title = item.Title
				db[i].Description = item.Description
				db[i].Code = item.Code
				db[i].Price = item.Price
	
				ctx.JSON(http.StatusOK, gin.H{"data": db[i]})
				return
			}
		}
	*/
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
	}

