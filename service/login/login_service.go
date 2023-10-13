package login

import "github.com/paula-michele-brisa/backend-campeonato/repository/login"

type loginService struct {
	loginRepository login.LoginRepositoryInterface
}
