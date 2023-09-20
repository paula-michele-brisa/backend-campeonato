package players

import "github.com/gin-gonic/gin"

// RegisterPlayerHandler efetua o registro/cadastro do jogador
func RegisterPlayerHandler(context *gin.Context) {
	context.JSON(201, gin.H{
		"message": "Jogar Cadastrado",
	})
}
