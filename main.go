package main

import (
	"EventManagementSystem/db"
	"EventManagementSystem/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // The server runs at port 8080
}
