package main

import (
	"database/sql"
	login3 "github.com/paula-michele-brisa/backend-campeonato/handler/login"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/repository/login"
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
	login2 "github.com/paula-michele-brisa/backend-campeonato/service/login"
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

func initLoginDependencies(database *sql.DB) login3.LoginHandlerInterface {
	loginRepository := login.LoginRepository(database)
	loginService := login2.LoginService(loginRepository)
	return login3.LoginHandler(loginService)
}
