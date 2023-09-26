package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"go.uber.org/zap"
)

func (user *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	user.EncryptPassword()
	return nil
}
