package main


import (
	"github.com/t-d-jesus/api-go-gin/routes"
	"github.com/t-d-jesus/api-go-gin/database"
	)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}