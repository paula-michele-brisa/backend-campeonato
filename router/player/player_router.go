package player

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/player"
)

func PLayerRouter(router *gin.RouterGroup) {

	playerGroup := router.Group("/player")

	// Listar jogadores cadastrados
	playerGroup.GET("/", player.FindPlayers)

	// Lista a quantidade de jogadores cadastrados
	playerGroup.GET("/countPlayer", player.FindTotalPlayer)

	// Cadastra um novo jogador
	playerGroup.POST("/", player.CreatePlayer)

	// Atualiza os dados do jogador
	playerGroup.PUT("/:id", player.UpdatePlayer)

	// Deleta jogador
	playerGroup.DELETE("/:id", player.DeletePlayer)

	// Lista jogador pelo ID
	playerGroup.GET("/:id", player.FindPlayerByID)

}
