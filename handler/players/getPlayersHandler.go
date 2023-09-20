package players

import "github.com/gin-gonic/gin"

// GetPlayers
func GetPlayersHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"jogador": "zezinho",
	})
}
