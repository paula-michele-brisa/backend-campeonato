package player

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/players"
)

// PlayerRouter rotas referentes a jogadores
func PlayerRouter(router *gin.RouterGroup) {
	playerGroup := router.Group("/team")

	// Obtem os jogadores cadastrados
	playerGroup.GET("/payers", players.GetPlayersHandler)

	// Cadastra um novo jogador
	playerGroup.POST("/player", players.RegisterPlayerHandler)

	// Editar/Atualizar jogador
	playerGroup.PUT("/player/:id", players.UpdatePlayerHandler)

	// Deleta um jogador
	playerGroup.DELETE("/player/:id", players.DeletePlayerHandler)
}
