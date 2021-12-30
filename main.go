package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/bookapi/models"
)

// let check db is working

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run()
}
