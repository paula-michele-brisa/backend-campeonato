package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
)

type TeamServiceInterface interface {
	CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr)
	FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr)
	FindTeamByEmail(email string) (team.TeamDomainInterface, *rest_err.RestErr)
	UpdateTeam(id string, team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr)
	DeleteTeam(id string) *rest_err.RestErr
	FindTotalRegisteredTeams() (int, *rest_err.RestErr)
	FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr)
}
