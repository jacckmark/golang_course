package main

import (
	"github.com/gin-gonic/gin"
	"rest.com/api/db"
	"rest.com/api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
