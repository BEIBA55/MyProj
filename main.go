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

	var defaultSubjects = []models.Subject{
		{Name: "Математика"},
		{Name: "Физика"},
		{Name: "Химия"},
		{Name: "История"},
		{Name: "Английский язык"},
	}

	for _, s := range defaultSubjects {
		var existing models.Subject
		result := config.DB.Where("name = ?", s.Name).First(&existing)
		if result.RowsAffected == 0 {
			config.DB.Create(&s)
		}
	}

	r := routes.SetupRouter()
	r.Run(":8080")
}
