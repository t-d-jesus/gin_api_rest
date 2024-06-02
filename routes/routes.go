package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/t-d-jesus/api-go-gin/controllers"
  )

func HandleRequests(){
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}