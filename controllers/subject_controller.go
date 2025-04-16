package controllers

import (
	"MyProj/config"
	"MyProj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	config.DB.Find(&subjects)
	c.JSON(http.StatusOK, subjects)
}

func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&subject)
	c.JSON(http.StatusOK, subject)
}

func UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	var subject models.Subject
	if err := config.DB.First(&subject, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&subject)
	c.JSON(http.StatusOK, subject)
}

func DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Subject{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted"})
}
