package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	fmt.Println("router == nil ? ", router == nil)
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
