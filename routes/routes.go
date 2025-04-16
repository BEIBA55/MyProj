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

	z := r.Group("/")
	z.Use(middleware.AuthMiddleware())

	// Students
	z.GET("/students", controllers.GetStudents)
	z.POST("/students", controllers.CreateStudent)
	z.PUT("/students/:id", controllers.UpdateStudent)
	z.DELETE("/students/:id", controllers.DeleteStudent)

	// Teachers
	z.GET("/teachers", controllers.GetTeachers)
	z.POST("/teachers", controllers.CreateTeacher)
	z.PUT("/teachers/:id", controllers.UpdateTeacher)
	z.DELETE("/teachers/:id", controllers.DeleteTeacher)

	z.GET("/subjects/:subject_id/teachers", controllers.GetTeachersBySubject)
	z.GET("/teachers/:teacher_id/students", controllers.GetStudentsByTeacher)

	z.GET("/subjects", controllers.GetSubjects)
	z.POST("/subjects", controllers.CreateSubject)
	z.PUT("/subjects/:id", controllers.UpdateSubject)
	z.DELETE("/subjects/:id", controllers.DeleteSubject)

	return r
}
