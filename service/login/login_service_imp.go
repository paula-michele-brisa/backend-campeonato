package login

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
	login2 "github.com/paula-michele-brisa/backend-campeonato/repository/login"
	"github.com/paula-michele-brisa/backend-campeonato/utils"
)

func (loginService *loginService) LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, string, *rest_err.RestErr) {

	utils.EncryptPasswordLogin(login)

	user, err := loginService.loginRepository.LoginRepository(login.GetEmail(), login.GetPassword())

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()

	if err != nil {

		return nil, "", err
	}

	println(token)

	return user, token, nil

}

func LoginService(loginRepository login2.LoginRepositoryInterface) LoginServiceInterface {
	return &loginService{
		loginRepository: loginRepository,
	}
}
