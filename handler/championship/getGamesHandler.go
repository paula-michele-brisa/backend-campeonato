package championship

import "github.com/gin-gonic/gin"

// GetGamesHandler obtém os jogos gerados no campeonato
func GetGamesHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"total": "GPS x 4 andar",
	})
}
