package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	"github.com/paula-michele-brisa/backend-campeonato/handler/user"
)

func routers(router *gin.Engine) {

	basePath := "api/v1"

	v1 := router.Group(basePath)
	{
		// Obter usuário
		v1.GET("/user", user.GetUserHandler)

		// Registrar usuário
		v1.POST("/user", user.RegisterUserHandler)

		// Efetuar login
		v1.POST("/login", user.LoginUserHandler)

		// Obter o total de times cadastrados
		v1.GET("/team_count", teams.GetTotalRegisteredTeamsHandler)

		// Listar os times cadastrados
		v1.GET("/teams", teams.GetTeamsHandler)

		// Cadastra um novo time
		v1.POST("/team", teams.RegisterTeamHandler)

		// Editar/Atualizar time
		v1.PUT("/team/:id", teams.UpdateTeamHandler)

		// Deleta um time
		v1.DELETE("/time/:id", teams.DeleteTeamHandler)

		// Obtem um time pelo id
		v1.GET("/team/:id", teams.GetTeamHandler)

		// Obtem os jogadores cadastrados
		v1.GET("/payers", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "10",
			})
		})

		// Cadastra um novo jogador
		v1.POST("/player", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Player Cadastrado",
			})
		})

		// Editar/Atualizar jogador
		v1.PUT("/player/:id", func(context *gin.Context) {
			context.JSON(400, gin.H{
				"message": "Jogador atualizado",
			})
		})

		// Deleta um time
		v1.DELETE("/player/:id", func(context *gin.Context) {
			context.JSON(204, gin.H{
				"message": "Jogador excluído",
			})
		})

		// Obtem os jogos do campeonato
		v1.GET("/games", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "GPS x 4 andar",
			})
		})

		// Gera os jogos
		v1.POST("/championship", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Campeonato gerado com sucesso!",
			})
		})

	}

}
