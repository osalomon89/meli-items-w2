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

func NewController(ucService domainPorts.UseCasesService, ucRepository domainPorts.Repository ) *controller{
	controllerService := controller{
		ucService: ucService,
		ucRepository: ucRepository,
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

	err = ctrl.ucRepository.UpdateItem(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  item,
	})
}


	func (ctrl *controller) GetItem(ctx *gin.Context) {
		request := ctx.Request
		body := request.Body
	
		var item dom.Item

		err := json.NewDecoder(body).Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data": err.Error(),
			})
			return
		}
		//Aca iria ucService o ucRepository?
		err = ctrl.ucService.GetItem(item)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": true,
				"data": err.Error(),
			})
		}

		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
	}