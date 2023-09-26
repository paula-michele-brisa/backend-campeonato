package user

import (
	"github.com/gin-gonic/gin"
	handler "github.com/paula-michele-brisa/backend-campeonato/handler/user"
)

// UserRouter contém as rotas referentes ao usuário
func UserRouter(router *gin.RouterGroup) {

	userGroup := router.Group("/user")

	// Obter usuário
	userGroup.GET("/", handler.GetUserHandler)

	// Registrar usuário
	userGroup.POST("/", handler.CreateUserHandler)

}
