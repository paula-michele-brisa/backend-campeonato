package main

import (
	"database/sql"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"

	repository2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
	"github.com/paula-michele-brisa/backend-campeonato/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)

	return user2.NewUserHandlerInterface(service)
}
