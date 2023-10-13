package main

import (
	"github.com/paula-michele-brisa/backend-campeonato/config"
	"github.com/paula-michele-brisa/backend-campeonato/router"
)

func main() {
	databasePostgres, err := config.NewDatabasePostgres()

	if err != nil {
		return
	}

	userHandler := initDependencies(databasePostgres)
	teamHandler := initTeamDependencies(databasePostgres)
	loginHandler := initLoginDependencies(databasePostgres)

	router.SetupRouter(userHandler, teamHandler, loginHandler)

}
