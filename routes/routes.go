package routes

import (
	"MyProj/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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

	return r
}
