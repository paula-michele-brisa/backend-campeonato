package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
)

type LoginRepositoryInterface interface {
	LoginRepository(login login.LoginDomainInterface) (login.LoginDomainInterface, *rest_err.RestErr)
}
