package team

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
)

// TeamRouter rotas para team
func TeamRouter(router *gin.RouterGroup) {
	teamGroup := router.Group("/team")

	// Obter o total de times cadastrados
	teamGroup.GET("/team_count", teams.GetTotalRegisteredTeamsHandler)

	// Listar os times cadastrados
	teamGroup.GET("/", teams.GetTeamsHandler)

	// Cadastra um novo time
	teamGroup.POST("/createTeam", teams.RegisterTeamHandler)

	// Editar/Atualizar time
	teamGroup.PUT("/updateTeam/:id", teams.UpdateTeamHandler)

	// Deleta um time
	teamGroup.DELETE("/deleteTeam/:id", teams.DeleteTeamHandler)

	// Obtem um time pelo id
	teamGroup.GET("/:id", teams.GetTeamHandler)
}
