package main

import (
	"github.com/api/routes"
	"github.com/gin-gonic/gin"
)
func main (){
	server := gin.Default()
	routes.SetupUserRoute(server)
	server.Run(":8080")
}