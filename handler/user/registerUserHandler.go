package user

import "github.com/gin-gonic/gin"

// RegisterUserHandler é responsável por retorar os dados do usuário
func RegisterUserHandler(context *gin.Context) {
	context.JSON(201, gin.H{
		"message": "Usuário cadastrado com sucesso!",
	})
}
