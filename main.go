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

	router.SetupRouter(userHandler)

}
