package controllers

import (
	"MyProj/config"
	"MyProj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.ShouldBindJSON(&student)
	config.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Student{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

func GetStudentsByTeacher(c *gin.Context) {
	teacherID := c.Param("teacher_id")
	var students []models.Student
	config.DB.Where("teacher_id = ?", teacherID).Find(&students)
	c.JSON(http.StatusOK, students)
}
