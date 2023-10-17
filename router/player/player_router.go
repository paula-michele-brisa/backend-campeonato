package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PLayerRouter(router *gin.RouterGroup) {

	playerGroup := router.Group("/player")

	// Listar jogadores cadastrados
	playerGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Jogadores ok!")

	})

	// Lista a quantidade de jogadores cadastrados
	playerGroup.GET("/countPlayer", func(context *gin.Context) {

	})

	// Cadastra um novo jogador
	playerGroup.POST("/createPlayer", func(context *gin.Context) {

	})

	// Atualiza os dados do jogador
	playerGroup.PUT("/:id", func(context *gin.Context) {

	})

	// Deleta jogador
	playerGroup.DELETE("/:id", func(context *gin.Context) {

	})

	// Lista jogador pelo ID
	playerGroup.GET("/:id", func(context *gin.Context) {

	})

}
