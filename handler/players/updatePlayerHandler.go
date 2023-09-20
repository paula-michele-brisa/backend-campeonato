package players

import "github.com/gin-gonic/gin"

// UpdatePlayerHandler atualiza
func UpdatePlayerHandler(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "Jogador atualizado",
	})
}
