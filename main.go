package main

import (
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {

	r := gin.Default()

	r.Run(port)

}
