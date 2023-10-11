package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/service/team"
)

func NewTeamHandlerInterface(teamService team.TeamDomainServiceInterface) TeamHandlerInterface {
	return &teamHandler{
		teamService: teamService,
	}

}

type TeamHandlerInterface interface {
	CreateTeamHandler(context *gin.Context)
	DeleteTeamHandler(context *gin.Context)
	FindTeamByIDHandler(context *gin.Context)
	FindTeams(context *gin.Context)
	FindTotalTeams(context *gin.Context)
	UpdateTeamHandler(context *gin.Context)
}

type teamHandler struct {
	teamService team.TeamDomainServiceInterface
}
