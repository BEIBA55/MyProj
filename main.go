package main

import (
	"MyProj/config"
	"MyProj/models"
	"MyProj/routes"
)

func main() {
	config.ConnectDB()

	// Автоматическая миграция моделей
	config.DB.AutoMigrate(&models.Subject{}, &models.Teacher{}, &models.Student{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
