package main

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
	team2 "github.com/paula-michele-brisa/backend-campeonato/service/team"

	repository2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
	"github.com/paula-michele-brisa/backend-campeonato/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)

	return user2.NewUserHandlerInterface(service)
}

func initTeamDependencies(database *sql.DB) teams.TeamHandlerInterface {
	teamRepository := team.TeamRepository(database)
	teamService := team2.TeamService(teamRepository)
	return teams.TeamHandler(teamService)

}
