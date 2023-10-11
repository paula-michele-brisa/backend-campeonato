package teams

import (
	"github.com/gin-gonic/gin"
	team2 "github.com/paula-michele-brisa/backend-campeonato/service/team"
)

// UpdateTeamHandler atualiza o time pelo id
func (team *teamHandler) UpdateTeamHandler(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "Time atualizado",
	})

}

// CreateTeamHandler cadastra um novo time
func (team *teamHandler) CreateTeamHandler(context *gin.Context) {

	context.JSON(201, gin.H{
		"message": "Time Cadastrado",
	})
}

// DeleteTeamHandler remove o time
func (team *teamHandler) DeleteTeamHandler(context *gin.Context) {
	context.JSON(204, gin.H{
		"message": "Time excluído",
	})
}

// GetTeamHandler obtém um time pelo id
func (team *teamHandler) FindTeamByIDHandler(context *gin.Context) {

	context.JSON(200, gin.H{
		"total": "10",
	})
}

// GetTeamsHandler retorna a lista de times cadastrados
func (team *teamHandler) FindTeams(context *gin.Context) {
	context.JSON(200, gin.H{
		"time": "GPS futebol clube",
	})

}

// GetTotalRegisteredTeamsHandler obtem a quantidade total de times cadastrados
func (team *teamHandler) FindTotalTeams(context *gin.Context) {
	context.JSON(200, gin.H{
		"total": "10",
	})
}

func TeamHandler(teamService team2.TeamServiceInterface) TeamHandlerInterface {
	return &teamHandler{
		teamService,
	}

}
