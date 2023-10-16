package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
	team2 "github.com/paula-michele-brisa/backend-campeonato/repository/team"
)

func (t *teamService) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {

	return nil, nil
}

func (t *teamService) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (t *teamService) FindTeamByEmail(email string) (team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (t *teamService) UpdateTeam(id string, team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (t *teamService) DeleteTeam(id string) *rest_err.RestErr {
	return nil
}

func (t *teamService) FindTotalRegisteredTeams() (int, *rest_err.RestErr) {
	return 0, nil
}

func (t *teamService) FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func TeamService(repository team2.TeamRepositoryInterface) TeamServiceInterface {
	return &teamService{
		repository: repository,
	}
}
