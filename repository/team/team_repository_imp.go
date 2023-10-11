package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/service/team"
)

func (team *teamRepository) CreateTeam(domainInterface team.TeamDomainServiceInterface) (team.TeamDomainServiceInterface, *rest_err.RestErr) {
	return nil, nil

}

func (team *teamRepository) FindTeamByID(id string) (team.TeamDomainServiceInterface, *rest_err.RestErr) {
	return nil, nil
}

func (team *teamRepository) FindTeamByEmail(email string) (team.TeamDomainServiceInterface, *rest_err.RestErr) {
	return nil, nil
}

func (team *teamRepository) UpdateTeam(id string, domainInterface team.TeamDomainServiceInterface) *rest_err.RestErr {
	return nil
}

func (team *teamRepository) DeleteTeam(string) *rest_err.RestErr {
	return nil
}
