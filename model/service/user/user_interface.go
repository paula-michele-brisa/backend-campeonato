package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	UpdateUser(string, user.UserDomainInterface) *rest_err.RestErr
	CreateUser(user.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*user.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
