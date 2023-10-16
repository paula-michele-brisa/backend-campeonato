package team

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
)

func TeamRepository(database *sql.DB) TeamRepositoryInterface {
	return &teamRepository{
		databaseConnection: database,
	}
}

func (t *teamRepository) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {

	return nil, nil
}

func (t *teamRepository) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (t *teamRepository) UpdateTeam(id string, team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (t *teamRepository) DeleteTeam(id string) *rest_err.RestErr {
	return nil
}

func (t *teamRepository) FindTotalRegisteredTeams() (int, *rest_err.RestErr) {
	return 0, nil
}

func (t *teamRepository) FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
