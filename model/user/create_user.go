package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
)

func (user *UserDomain) CreateUser() *rest_err.RestErr {
	user.EncryptPassword()
	return nil
}
