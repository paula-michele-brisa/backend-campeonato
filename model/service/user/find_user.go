package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"go.uber.org/zap"
)

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
