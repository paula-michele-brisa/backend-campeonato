package team

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
	team2 "github.com/paula-michele-brisa/backend-campeonato/entity/team"
	team3 "github.com/paula-michele-brisa/backend-campeonato/repository/convert/team"
)

func TeamRepository(database *sql.DB) TeamRepositoryInterface {
	return &teamRepository{
		databaseConnection: database,
	}
}

func (t *teamRepository) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {

	db := t.databaseConnection

	value := team3.ConvertDomainToEntity(team)

	query := `INSERT INTO t_team(name, badgePhoto, city, coach, website)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(query, value.Name, value.BadgePhoto, value.City, value.Coach, value.Website).Scan(&value.ID)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return team3.ConvertTeamEntityToDomain(*value), nil
}

func (t *teamRepository) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	db := t.databaseConnection

	query := `SELECT * from t_team where id=$1`

	teamEntity := &team2.TeamEntity{}

	err := db.QueryRow(query, id).Scan(&teamEntity.ID, &teamEntity.Name, &teamEntity.City, &teamEntity.Coach, &teamEntity.Website, &teamEntity.BadgePhoto)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return team3.ConvertTeamEntityToDomain(*teamEntity), nil
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
