package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"go.uber.org/zap"
)

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
