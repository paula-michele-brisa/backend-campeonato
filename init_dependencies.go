package main

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
	repository2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
	team2 "github.com/paula-michele-brisa/backend-campeonato/service/team"
	"github.com/paula-michele-brisa/backend-campeonato/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)
	return user2.NewUserHandlerInterface(service)
}

func initTeamDependencies(database *sql.DB) teams.TeamHandlerInterface {
	repository := team.TeamRespository(database)
	service := team2.TeamDomainService(repository)
	return teams.NewTeamHandlerInterface(service)
}
