package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
)

func (login *loginRepository) LoginRepository(email, password string) (login.LoginDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
