package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/championship"
	"github.com/paula-michele-brisa/backend-campeonato/handler/players"
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
		v1.GET("/payers", players.GetPlayersHandler)

		// Cadastra um novo jogador
		v1.POST("/player", players.RegisterPlayerHandler)

		// Editar/Atualizar jogador
		v1.PUT("/player/:id", players.UpdatePlayerHandler)

		// Deleta um jogador
		v1.DELETE("/player/:id", players.DeletePlayerHandler)

		// Obtem os jogos do campeonato
		v1.GET("/games", championship.GetGamesHandler)

		// Gera os jogos
		v1.POST("/championship", championship.GenerateChampionship)

	}

}
