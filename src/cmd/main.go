package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Erick-Bueno/go-api/routes"
)
func main (){
	server := gin.Default()
	routes.SetupUserRoute(server)
	server.Run(":8080")
}