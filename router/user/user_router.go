package user

import (
	"github.com/gin-gonic/gin"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
)

// UserRouters rotas usuário
func UserRouters(router *gin.RouterGroup, userHandler user2.UserHandlerInterface) {

	routerGroup := router.Group("/user")

	routerGroup.GET("/", userHandler.GetUserHandler)

	routerGroup.POST("/", userHandler.CreateUserHandler)

	routerGroup.PUT("/", userHandler.UpdateUserHandler)

	routerGroup.DELETE("/", userHandler.DeleteUserHandler)

}
