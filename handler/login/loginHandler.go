package login

import "github.com/gin-gonic/gin"

// LoginUserHandler efetua a conexão do usuário com o sistema
func LoginUserHandler(context *gin.Context) {
	context.JSON(201, gin.H{
		"message": "Usuário logado com sucesso!",
	})

}
