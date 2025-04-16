package main

import (
	"MyProj/config"
	"MyProj/models"
	"MyProj/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Subject{}, &models.Teacher{}, &models.Student{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
