package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/t-d-jesus/api-go-gin/models"
	"github.com/t-d-jesus/api-go-gin/database"
	)

func ShowStudents(c *gin.Context){
	c.JSON(http.StatusOK, models.Students)
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