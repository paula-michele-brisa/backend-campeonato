package team

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
)

func (t *teamDomainService) CreateTeamService(team TeamDomainServiceInterface) (TeamDomainServiceInterface, *rest_err.RestErr) {
	team, err := t.teamRespository.CreateTeam(team)

	if err != nil {
		logger.Error("Ocorreu um erro ao chamar o repository para criar time", err)
		return nil, err
	}

	return team, nil
}

func (team *teamDomainService) DeleteTeamService(context *gin.Context) {}

func (team *teamDomainService) FindTeamByIDService(context *gin.Context) {}

func (team *teamDomainService) FindTeamsService(context *gin.Context) {}

func (team *teamDomainService) FindTotalTeamsService(context *gin.Context) {}

func (team *teamDomainService) UpdateTeamService(context *gin.Context) {}

func TeamDomainService(teamRespository team.TeamRepositoryInterface) TeamDomainServiceInterface {
	return &teamDomainService{
		teamRespository: teamRespository,
	}
}
