package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	user2 "github.com/paula-michele-brisa/backend-campeonato/model/repository/user"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func NewUserDomainService(userRepository user2.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository user2.UserRepository
}

type UserDomainService interface {
	UpdateUser(string, user.UserDomainInterface) *rest_err.RestErr
	CreateUser(user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
