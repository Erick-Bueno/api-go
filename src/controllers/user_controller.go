package controllers

import(
	"net/http"
	"github.com/api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


//server pra definirmos as dependencias que vamos ter no nosso controller como por exemplo os services, e o equivalente a nossa classe controller em poo
type userController struct{
}
//essa funcao representa o construtor que no caso seria o construtor da nossa classe controller se estivessemos em poo
func NewUserController() userController {
	return userController{}
}
//representa o metodo da nossa classe, nesse caso get producs
func (u *userController) GetUsers(ctx *gin.Context){
	users := []model.User{
		{
			ID: 1,
			ExternalId: uuid.New(),
			Name: "erick",
			Age: 21,
		},
	}
	ctx.JSON(http.StatusOK, users)
}

