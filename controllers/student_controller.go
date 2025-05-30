package controllers

import (
	"MyProj/config"
	"MyProj/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStudents(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	fmt.Printf("Received page: %d, limit: %d\n", page, limit)

	offset := (page - 1) * limit

	var students []models.Student
	var totalStudents int64

	config.DB.Model(&models.Student{}).Count(&totalStudents)

	config.DB.Offset(offset).Limit(limit).Find(&students)

	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"total":      totalStudents,
		"totalPages": (totalStudents + int64(limit) - 1) / int64(limit),
		"students":   students,
	})
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
	var students []models.Student
	teacherID := c.DefaultQuery("teacher_id", "0")
	config.DB.Select("name").Where("teacher_id = ?", teacherID).Order("name").Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentCount(c *gin.Context) {
	var count int64
	config.DB.Model(&models.Student{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"student_count": count})
}
func GetStudentsWithoutTeacher(c *gin.Context) {
	var students []models.Student
	if err := config.DB.Where("teacher_id IS NULL").Find(&students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to retrieve students without teacher"})
		return
	}

	c.JSON(http.StatusOK, students)
}
