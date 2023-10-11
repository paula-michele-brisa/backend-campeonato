package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
)

type UserDomainService interface {
	UpdateUser(string, user.UserDomainInterface) *rest_err.RestErr
	CreateUser(user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
