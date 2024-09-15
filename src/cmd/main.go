package main

import (
	"github.com/Erick-Bueno/go-api/routes"
	"github.com/gin-gonic/gin"
)
func main (){
	server := gin.Default()
	routes.SetupUserRoute(server)
	server.Run(":8080")
}