package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"github.com/paula-michele-brisa/backend-campeonato/model/utils"
	"go.uber.org/zap"
)

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
