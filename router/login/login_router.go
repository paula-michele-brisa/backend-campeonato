package login

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/login"
)

func LoginRouter(router *gin.RouterGroup, login login.LoginHandlerInterface) {

	loginGroup := router.Group("/login")

	loginGroup.GET("/", login.LoginHandler)

}
