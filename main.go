package main

import (
	"go-01/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-01/configs"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	router := gin.Default()

	routes.PersonRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	router.Run()
}
