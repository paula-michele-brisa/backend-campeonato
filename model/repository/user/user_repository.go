package user

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func NewUserRespository(database *sql.DB) UserRepository {
	return &userRespository{
		database,
	}
}

type userRespository struct {
	databaseConnection *sql.DB
}

type UserRepository interface {
	CreateUser(domainInterface user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (user.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, domainInterface user.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
