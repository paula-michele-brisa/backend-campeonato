package players

import "github.com/gin-gonic/gin"

// DeletePlayerHandler remover o jogador
func DeletePlayerHandler(context *gin.Context) {
	context.JSON(204, gin.H{
		"message": "Jogador excluído",
	})
}
