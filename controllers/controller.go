package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/t-d-jesus/api-go-gin/models"
	)

func ShowStudents(c *gin.Context){
	c.JSON(http.StatusOK, models.Students)
}