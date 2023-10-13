package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
)

func (loginService *loginService) LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, *rest_err.RestErr) {
	loginDomain, err := loginService.loginRepository.LoginRepository(login)

	if err != nil {
		return nil, err
	}

	return loginDomain, nil

}
