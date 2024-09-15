package routes

import (

	"github.com/Erick-Bueno/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoute(server *gin.Engine){
	userController := controllers.NewUserController()
	userGroup := server.Group("/users")
	{
		userGroup.GET("/", userController.GetUsers)
	}
}