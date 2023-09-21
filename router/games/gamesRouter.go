package games

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/championship"
)

// GamesRouter rotas referentes aos jogos e geração de campeonato
func GamesRouter(router *gin.RouterGroup) {

	gamersGroup := router.Group("/games")

	// Obtem os jogos do campeonato
	gamersGroup.GET("/games", championship.GetGamesHandler)

	// Gera os jogos
	gamersGroup.POST("/championship", championship.GenerateChampionship)
}
