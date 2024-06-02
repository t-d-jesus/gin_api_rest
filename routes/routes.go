package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/t-d-jesus/api-go-gin/controllers"
  )

func HandleRequests(){
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.ShowStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}