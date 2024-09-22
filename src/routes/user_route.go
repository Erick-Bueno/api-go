package routes

import (
	"github.com/Erick-Bueno/go-api/controllers"
	"github.com/Erick-Bueno/go-api/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoute(server *gin.Engine){
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)
	userGroup := server.Group("/users")
	{
		userGroup.GET("/", userController.GetUsers)
	}
}