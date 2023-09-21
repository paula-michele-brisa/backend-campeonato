package login

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/login"
)

// LoginRouter efetua login do usuário
func LoginRouter(router *gin.RouterGroup) {

	loginGroup := router.Group("/login")

	// Efetuar login
	loginGroup.POST("/login", login.LoginUserHandler)
}
