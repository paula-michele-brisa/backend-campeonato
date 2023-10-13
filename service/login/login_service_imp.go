package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
	login2 "github.com/paula-michele-brisa/backend-campeonato/repository/login"
	"github.com/paula-michele-brisa/backend-campeonato/utils"
)

func (loginService *loginService) LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, *rest_err.RestErr) {

	utils.EncryptPasswordLogin(login)

	loginDomain, err := loginService.loginRepository.LoginRepository(login.GetEmail(), login.GetPassword())

	if err != nil {
		return nil, err
	}

	return loginDomain, nil

}

func LoginService(loginRepository login2.LoginRepositoryInterface) LoginServiceInterface {
	return &loginService{
		loginRepository: loginRepository,
	}
}
