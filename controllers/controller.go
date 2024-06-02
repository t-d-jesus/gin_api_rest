package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

func ShowStudents(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
	  "id": "1",
	  "name":"Joseph",
	})
}

