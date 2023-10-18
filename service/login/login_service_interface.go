package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
)

type LoginServiceInterface interface {
	LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, string, *rest_err.RestErr)
}
