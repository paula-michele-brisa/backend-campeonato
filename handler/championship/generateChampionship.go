package championship

import "github.com/gin-gonic/gin"

// GenerateChampionship gera os jogos do campeonato
func GenerateChampionship(context *gin.Context) {
	context.JSON(201, gin.H{
		"message": "Campeonato gerado com sucesso!",
	})
}
