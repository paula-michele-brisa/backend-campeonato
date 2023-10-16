package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
	team2 "github.com/paula-michele-brisa/backend-campeonato/repository/team"
)

func (t *teamService) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {

	teamDomain, err := t.repository.CreateTeam(team)

	if err != nil {
		logger.Error("Erro ao tentar criar time", err)
		return nil, err
	}

	return teamDomain, nil
}

func (t *teamService) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	teamDomain, err := t.repository.FindTeamByID(id)

	if err != nil {
		logger.Error("Erro ao tentar encontrar time", err)
		return nil, err
	}

	return teamDomain, nil
}

func (t *teamService) UpdateTeam(id string, team team.TeamDomainInterface) *rest_err.RestErr {
	err := t.repository.UpdateTeam(id, team)

	if err != nil {
		logger.Error("Erro ao tentar atualizar time", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

func (t *teamService) DeleteTeam(id string) *rest_err.RestErr {
	err := t.repository.DeleteTeam(id)

	if err != nil {
		logger.Error("Erro ao tentar atualizar time", err)
		return err
	}

	return nil
}

func (t *teamService) FindTotalRegisteredTeams() (int, *rest_err.RestErr) {
	totalTeams, err := t.repository.FindTotalRegisteredTeams()

	if err != nil {
		logger.Error("Erro ao tentar atualizar time", err)
		return 0, err
	}

	return totalTeams, nil
}

func (t *teamService) FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr) {
	totalTeams, err := t.repository.FindTotalTeams()

	if err != nil {
		logger.Error("Erro ao tentar atualizar time", err)
		return nil, err
	}

	return totalTeams, nil
}

func TeamService(repository team2.TeamRepositoryInterface) TeamServiceInterface {
	return &teamService{
		repository: repository,
	}
}
