package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	r := gin.Default()
	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, []map[string]interface{}{
			{"id": 1, "name": "Student 1", "age": 20},
			{"id": 2, "name": "Student 2", "age": 22},
		})
	})

	req, _ := http.NewRequest("GET", "/students?page=1&limit=10", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Student 1")
	assert.Contains(t, w.Body.String(), "Student 2")
}

func TestCreateStudent(t *testing.T) {
	r := gin.Default()
	r.POST("/students", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"id": 1, "name": "New Student", "age": 21})
	})

	req, _ := http.NewRequest("POST", "/students", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Student")
}

func TestUpdateStudent(t *testing.T) {
	r := gin.Default()
	r.PUT("/students/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": 1, "name": "Updated Student", "age": 23})
	})

	reqBody := strings.NewReader(`{"name": "Updated Student", "age": 23}`)
	req, _ := http.NewRequest("PUT", "/students/1", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Student")
}

func TestDeleteStudent(t *testing.T) {
	r := gin.Default()
	r.DELETE("/students/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	})

	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Student deleted")
}

func TestGetTeachers(t *testing.T) {
	r := gin.Default()
	r.GET("/teachers", func(c *gin.Context) {
		c.JSON(http.StatusOK, []map[string]interface{}{
			{"id": 1, "name": "Teacher 1"},
			{"id": 2, "name": "Teacher 2"},
		})
	})

	req, _ := http.NewRequest("GET", "/teachers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Teacher 1")
	assert.Contains(t, w.Body.String(), "Teacher 2")
}

func TestCreateTeacher(t *testing.T) {
	r := gin.Default()
	r.POST("/teachers", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"id": 3, "name": "New Teacher"})
	})

	req, _ := http.NewRequest("POST", "/teachers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Teacher")
}

func TestUpdateTeacher(t *testing.T) {
	r := gin.Default()
	r.PUT("/teachers/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": 1, "name": "Updated Teacher"})
	})

	reqBody := strings.NewReader(`{"name": "Updated Teacher"}`)
	req, _ := http.NewRequest("PUT", "/teachers/1", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Teacher")
}

func TestDeleteTeacher(t *testing.T) {
	r := gin.Default()
	r.DELETE("/teachers/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
	})

	req, _ := http.NewRequest("DELETE", "/teachers/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Teacher deleted")
}

func TestFilterStudentsByTeacherID(t *testing.T) {
	r := gin.Default()
	r.GET("/students", func(c *gin.Context) {
		teacherID := c.Query("teacher_id")
		c.JSON(http.StatusOK, gin.H{"filtered_by_teacher": teacherID})
	})

	req, _ := http.NewRequest("GET", "/students?teacher_id=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "2")
}

func TestPaginationLimit(t *testing.T) {
	r := gin.Default()
	r.GET("/students", func(c *gin.Context) {
		limit := c.Query("limit")
		c.JSON(http.StatusOK, gin.H{"limit": limit})
	})

	req, _ := http.NewRequest("GET", "/students?limit=5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "5")
}
