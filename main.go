package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/bookapi/controller"
	"github.com/krishpranav/bookapi/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	/* routes */
	r.GET("/books", controller.FindBooks)

	r.Run()
}
