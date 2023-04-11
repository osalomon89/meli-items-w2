package controller

import (
	"apiRestPractice/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type responseInfo struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

var data []model.Item

func DataMock() {
	jsonFile, errorJson := os.Open("assets/items.json")
	if errorJson != nil {
		fmt.Printf("errorJson: %s", errorJson.Error())
	}
	defer jsonFile.Close()

	byteValue, errorByte := io.ReadAll(jsonFile)
	if errorByte != nil {
		fmt.Printf("errorByte: %s", errorByte.Error())
	}

	var result []model.Item
	json.Unmarshal(byteValue, &result)

	data = result
}

func checkFields(req map[string]any) []error {
	var errors []error
	if req["code"] == "" || req["code"] == nil {
		errors = append(errors, fmt.Errorf("code is required"))
	}
	if req["title"] == "" || req["title"] == nil {
		errors = append(errors, fmt.Errorf("title is required"))
	}
	if req["description"] == "" || req["description"] == nil {
		errors = append(errors, fmt.Errorf("description is required"))
	}
	if req["price"] == "" || req["price"] == nil {
		errors = append(errors, fmt.Errorf("price is required"))
	}
	if req["stock"] == "" || req["stock"] == nil {
		errors = append(errors, fmt.Errorf("stock is required"))
	}
	return errors
}

func newItem(result map[string]any) (item model.Item) {
	location, errLo := time.LoadLocation("Africa/Conakry")
	if errLo != nil {
		fmt.Printf("Error in location: %s", errLo.Error())
	}
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	item.Code = fmt.Sprint(result["code"])
	item.Title = fmt.Sprint(result["title"])
	item.Description = fmt.Sprint(result["description"])
	item.Price, _ = strconv.Atoi(fmt.Sprint(result["price"]))
	item.Stock, _ = strconv.Atoi(fmt.Sprint(result["stock"]))
	item.Id = len(data) + 10
	dateStrg := strings.Split(currentTime, " ")
	item.CreatedAt = dateStrg[0] + "T" + dateStrg[1] + "Z"
	item.UpdatedAt = item.CreatedAt
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
	return item
}

func updateItem(result map[string]any, itemToUpdate model.Item) (item model.Item) {
	location, errLo := time.LoadLocation("Africa/Conakry")
	if errLo != nil {
		fmt.Printf("Error in location: %s", errLo.Error())
	}
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	item.Code = fmt.Sprint(result["code"])
	item.Title = fmt.Sprint(result["title"])
	item.Description = fmt.Sprint(result["description"])
	item.Price, _ = strconv.Atoi(fmt.Sprint(result["price"]))
	item.Stock, _ = strconv.Atoi(fmt.Sprint(result["stock"]))
	item.Id = itemToUpdate.Id
	dateStrg := strings.Split(currentTime, " ")
	item.CreatedAt = itemToUpdate.CreatedAt
	item.UpdatedAt = dateStrg[0] + "T" + dateStrg[1] + "Z"
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
	return item
}

func AddItem(c *gin.Context) {
	//----------------
	req := c.Request
	body := req.Body
	var result map[string]any
	resBody, _ := io.ReadAll(body)
	//err := json.NewDecoder(body).Decode(&item)
	err := json.Unmarshal(resBody, &result)
	if len(checkFields(result)) > 0 {
		var errToString string
		for _, error := range checkFields(result) {
			errToString = errToString + fmt.Sprintf("%s, ", error.Error())
		}
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid body struct : %s", errToString),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid json: %s", err.Error()),
		})
		return
	}
	//----------------
	for _, item := range data {
		if item.Code == fmt.Sprint(result["code"]) {
			c.JSON(http.StatusBadRequest, responseInfo{
				Error: true,
				Data:  fmt.Sprint("Invalid item, code is already in database"),
			})
			return
		}
	}
	fmt.Println("Guardando el item...")
	data = append(data, newItem(result))
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  fmt.Sprintf("Item saved!"),
	})
	return
}

func UpdateItem(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	//----------------
	req := c.Request
	body := req.Body
	var result map[string]any
	resBody, _ := io.ReadAll(body)
	//err := json.NewDecoder(body).Decode(&item)
	err := json.Unmarshal(resBody, &result)
	if len(checkFields(result)) > 0 {
		var errToString string
		for _, error := range checkFields(result) {
			errToString = errToString + fmt.Sprintf("%s, ", error.Error())
		}
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid body struct : %s", errToString),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid json: %s", err.Error()),
		})
		return
	}
	//----------------
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
	}
	var idFoundFlag = true
	for i, item := range data {
		if id == item.Id {
			idFoundFlag = false
			fmt.Println("Modificando el item...")
			data[i] = updateItem(result, item)
			fmt.Println(data)
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  fmt.Sprintf("Item updated!"),
			})
			return
		}
	}
	if idFoundFlag {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprint("Id not found"),
		})
	}
}

func GetById(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
	}

	var idFoundFlag = true
	for _, item := range data {
		if item.Id == id {
			idFoundFlag = false
			itemToString, errMarshal := json.Marshal(item)
			if errMarshal != nil {
				c.JSON(http.StatusInternalServerError, responseInfo{
					Error: true,
					Data:  fmt.Sprintf("Error marshal item : %s", errMarshal.Error()),
				})
				return
			}
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  string(itemToString),
			})
			return
		}
	}
	if idFoundFlag {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprint("Id not found"),
		})
		return
	}
}

func DeleteById(c *gin.Context) {
	id, errParseInt := strconv.Atoi(c.Param("id"))
	if errParseInt != nil {
		fmt.Sprintf("Error parse int: %s", errParseInt.Error())
	}

	var idFoundFlag = true
	for index, item := range data {
		if item.Id == id {
			idFoundFlag = false
			data = append(data[:index], data[index+1:]...)
			c.JSON(http.StatusOK, responseInfo{
				Error: false,
				Data:  fmt.Sprint("Item deleted"),
			})
			fmt.Println(data)
			return
		}
	}
	if idFoundFlag {
		c.JSON(http.StatusNotFound, responseInfo{
			Error: true,
			Data:  fmt.Sprint("Id not found"),
		})
		return
	}
}

func GetByStatusAndLimit(c *gin.Context) {
	statusReq := c.Query("status")
	limitReq := c.Query("limit")
	var itemsRes []model.Item
	//sort.Slice(data, func(i, j int) bool {
	//	dateI, errorI := time.Parse("2006-01-02", dateString)
	//
	//	if errorI != nil {
	//		fmt.Println(errorI)
	//		//return
	//	}
	//})
	limitInt, _ := strconv.Atoi(limitReq)
	if limitInt == 0 && limitReq == "" {
		limitInt = 10
	}
	if limitInt > 20 {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprint("The maximum limit value is 20"),
		})
		return
	}
	fmt.Println(limitInt)
	if statusReq == "ACTIVE" || statusReq == "INACTIVE" {
		for _, item := range data {
			if statusReq == item.Status && len(itemsRes) < limitInt {
				itemsRes = append(itemsRes, item)
			}
		}
	} else if statusReq == "" {
		for _, item := range data {
			if len(itemsRes) < limitInt {
				itemsRes = append(itemsRes, item)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, responseInfo{
			Error: true,
			Data:  fmt.Sprintf("Invalid status : %s", statusReq),
		})
		return
	}
	response := make(map[string]any)
	response["totalPages"] = len(itemsRes)
	response["data"] = itemsRes
	c.JSON(http.StatusOK, responseInfo{
		Error: false,
		Data:  response,
	})
	return
}

func dateToInt() {

}
