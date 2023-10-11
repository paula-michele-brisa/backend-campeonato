package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/service/team"
)

type TeamRepositoryInterface interface {
	CreateTeam(domainInterface team.TeamDomainServiceInterface) (team.TeamDomainServiceInterface, *rest_err.RestErr)
	FindTeamByID(id string) (team.TeamDomainServiceInterface, *rest_err.RestErr)
	FindTeamByEmail(email string) (team.TeamDomainServiceInterface, *rest_err.RestErr)
	UpdateTeam(id string, domainInterface team.TeamDomainServiceInterface) *rest_err.RestErr
	DeleteTeam(string) *rest_err.RestErr
}
