package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	user2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
)

func NewUserDomainService(userRepository user2.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

func (user *userDomainService) CreateUser(
	userDomain user.UserDomainInterface,
) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.userRepository.CreateUser(userDomain)

}

func (user *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	return user.userRepository.DeleteUser(userId)

}

func (user *userDomainService) FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.userRepository.FindUserByEmail(email)
}

func (user *userDomainService) FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.userRepository.FindUserByID(id)
}

func (user *userDomainService) UpdateUser(id string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	return user.userRepository.UpdateUser(id, userDomain)

}
