package routes

import (
	"MyProj/controllers"
	"MyProj/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Students
	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	// Teachers
	r.GET("/teachers", controllers.GetTeachers)
	r.POST("/teachers", controllers.CreateTeacher)
	r.PUT("/teachers/:id", controllers.UpdateTeacher)
	r.DELETE("/teachers/:id", controllers.DeleteTeacher)

	r.GET("/subjects/:subject_id/teachers", controllers.GetTeachersBySubject)
	r.GET("/teachers/:teacher_id/students", controllers.GetStudentsByTeacher)

	r.GET("/subjects", controllers.GetSubjects)
	r.POST("/subjects", controllers.CreateSubject)
	r.PUT("/subjects/:id", controllers.UpdateSubject)
	r.DELETE("/subjects/:id", controllers.DeleteSubject)

	return r
}
