package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
	team2 "github.com/paula-michele-brisa/backend-campeonato/repository/team"
)

func (t *teamService) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {
	return t.repository.CreateTeam(team)

}

func (t *teamService) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	return t.repository.FindTeamByID(id)
}

func (t *teamService) UpdateTeam(id string, team team.TeamDomainInterface) *rest_err.RestErr {
	return t.repository.UpdateTeam(id, team)
}

func (t *teamService) DeleteTeam(id string) *rest_err.RestErr {
	return t.repository.DeleteTeam(id)
}

func (t *teamService) FindTotalRegisteredTeams() (int, *rest_err.RestErr) {
	return t.repository.FindTotalRegisteredTeams()
}

func (t *teamService) FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr) {
	return t.repository.FindTotalTeams()
}

func TeamService(repository team2.TeamRepositoryInterface) TeamServiceInterface {
	return &teamService{
		repository: repository,
	}
}
