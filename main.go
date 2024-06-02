package main


import (
	"github.com/t-d-jesus/api-go-gin/routes"
	"github.com/t-d-jesus/api-go-gin/models"
	)

func main() {
	models.Students = []models.Student{
		{Name: "Rau Raiz", CPF: "11111111120", RG: "222222223"},
		{Name: "Ana Laura", CPF: "11111111130", RG: "222222224"},		
	}
	routes.HandleRequests()
}