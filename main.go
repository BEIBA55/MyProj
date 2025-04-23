package main

import (
	"MyProj/config"
	"MyProj/routes"
)

func main() {
	config.ConnectDB()

	//config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
