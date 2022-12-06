package main

import (
	"github.com/farisamirmudin/go-crud/initializers"
	"github.com/farisamirmudin/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
