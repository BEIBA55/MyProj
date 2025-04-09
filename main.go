package main

import (
	"MyProj/config"
	"MyProj/models"
	"MyProj/routes"
)

func main() {
	config.ConnectDB()

	// Автоматическая миграция моделей
	config.DB.AutoMigrate(&models.Student{}, &models.Teacher{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
