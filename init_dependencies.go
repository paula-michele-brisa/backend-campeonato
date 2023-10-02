package main

import (
	"database/sql"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	repository2 "github.com/paula-michele-brisa/backend-campeonato/model/repository"
	"github.com/paula-michele-brisa/backend-campeonato/model/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)
	return user2.NewUserHandlerInterface(service)
}
