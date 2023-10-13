package teams

import (
	"github.com/gin-gonic/gin"
	team2 "github.com/paula-michele-brisa/backend-campeonato/domain/team"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/view"
	"net/http"
)

// UpdateTeamHandler atualiza o time pelo id
func (team *teamHandler) UpdateTeamHandler(context *gin.Context) {

	var teamRequest request.TeamRequest

	teamID := context.Param("teamID")

	if err := context.ShouldBindJSON(teamRequest); err != nil {

		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar dados do time")
		return
	}

	teamDomain := team2.NewTeamDomain(
		teamRequest.Name,
		teamRequest.Coach,
		teamRequest.Website,
		teamRequest.City,
		teamRequest.BadgePhoto,
	)

	teamResponse, err := team.teamService.UpdateTeam(teamID, teamDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	context.JSON(http.StatusOK, view.ConvertTeamDomainToResponse(teamResponse))

}

// CreateTeamHandler cadastra um novo time
func (team *teamHandler) CreateTeamHandler(context *gin.Context) {

	var teamRequest request.TeamRequest

	var totalRegisteredTeams = 8

	if err := context.ShouldBindJSON(teamRequest); err != nil {

		context.JSON(http.StatusBadRequest, "Erro ao tentar criar dados do time")
		return
	}

	totalTeams, err := team.teamService.FindTotalRegisteredTeams()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	if totalTeams > totalRegisteredTeams {
		context.JSON(http.StatusBadRequest, "O limite total de cadastro de times foi atingido.")
		return

	}

	teamDomain := team2.NewTeamDomain(
		teamRequest.Name,
		teamRequest.Coach,
		teamRequest.Website,
		teamRequest.City,
		teamRequest.BadgePhoto,
	)

	teamResponse, err := team.teamService.CreateTeam(teamDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	context.JSON(http.StatusOK, view.ConvertTeamDomainToResponse(teamResponse))
}

// DeleteTeamHandler remove o time
func (team *teamHandler) DeleteTeamHandler(context *gin.Context) {
	var teamRequest request.TeamRequest

	teamID := context.Param("teamID")

	if err := context.ShouldBindJSON(teamRequest); err != nil {

		context.JSON(http.StatusBadRequest, "Erro ao tentar deletar time")
		return
	}

	if err := team.teamService.DeleteTeam(teamID); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar deletar time.")
		return

	}

	context.JSON(http.StatusOK, "Time excluído com sucesso!")
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
