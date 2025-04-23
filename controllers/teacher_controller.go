package controllers

import (
	"MyProj/config"
	"MyProj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	config.DB.Find(&teachers)
	c.JSON(http.StatusOK, teachers)
}

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&teacher)

	config.DB.Preload("Subject").First(&teacher, teacher.ID)

	c.JSON(http.StatusOK, teacher)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := config.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&teacher)
	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Teacher{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
}

func GetTeachersBySubject(c *gin.Context) {
	subjectID := c.Param("subject_id")
	var teachers []models.Teacher
	config.DB.Where("subject_id = ?", subjectID).Find(&teachers)
	c.JSON(http.StatusOK, teachers)
}

func GetTeacherCount(c *gin.Context) {
	var count int64
	config.DB.Model(&models.Teacher{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"teacher_count": count})
}
