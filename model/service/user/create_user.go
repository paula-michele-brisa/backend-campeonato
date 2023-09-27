package user

import (
	"fmt"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"go.uber.org/zap"
)

func (user *userDomainService) CreateUser(
	userDomain user.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
