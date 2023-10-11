package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	user2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
	"github.com/paula-michele-brisa/backend-campeonato/utils"
	"go.uber.org/zap"
)

func NewUserDomainService(userRepository user2.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

func (user *userDomainService) CreateUser(
	userDomain user.UserDomainInterface,
) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	utils.EncryptPassword(userDomain)

	logger.Info(userDomain.GetPassword())

	userDomainRepository, err := user.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}
	return userDomainRepository, nil

}

func (user *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	err := user.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (user *userDomainService) FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("Journey", "findUserByEmail"))
	return user.userRepository.FindUserByEmail(email)
}

func (user *userDomainService) FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("Journey", "findUserByID"))
	return user.userRepository.FindUserByID(id)
}

func (user *userDomainService) UpdateUser(id string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser handler",
		zap.String("journey", "createUser"))

	logger.Info(userDomain.GetPassword())
	err := user.userRepository.UpdateUser(id, userDomain)

	if err != nil {
		return err
	}

	return nil
}
