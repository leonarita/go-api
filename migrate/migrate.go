package main

import (
	"go-01/configs"
	"go-01/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
