package main

import (
	"go_curd/initializers"
	"go_curd/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
		
}

func main() {
	initializers.DB.AutoMigrate(&models.Test{})
}