package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/t-d-jesus/api-go-gin/models"
	"github.com/t-d-jesus/api-go-gin/database"
	)

func ShowStudents(c *gin.Context){
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context){
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return	
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)

}

func ShowStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}	

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted student"})
}