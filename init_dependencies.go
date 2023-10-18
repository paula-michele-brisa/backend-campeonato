package main

import (
	"database/sql"
	login3 "github.com/paula-michele-brisa/backend-campeonato/handler/login"
	"github.com/paula-michele-brisa/backend-campeonato/handler/player"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/repository/login"
	player2 "github.com/paula-michele-brisa/backend-campeonato/repository/player"
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
	login2 "github.com/paula-michele-brisa/backend-campeonato/service/login"
	player3 "github.com/paula-michele-brisa/backend-campeonato/service/player"
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
	return teams.NewTeamHandler(teamService)

}

func initLoginDependencies(database *sql.DB) login3.LoginHandlerInterface {
	loginRepository := login.LoginRepository(database)
	loginService := login2.LoginService(loginRepository)
	return login3.LoginHandler(loginService)
}

func initPlayerDependencies(database *sql.DB) player.PlayerHandlerInterface {
	playerRepository := player2.PlayerRepository(database)
	playerService := player3.NewPlayerService(playerRepository)
	playerHandler := player.NewPlayerHandler(playerService)

	return playerHandler
}
