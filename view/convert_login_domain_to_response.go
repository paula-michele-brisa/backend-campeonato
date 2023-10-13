package view

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/response"
)

func ConverteToLoginDomainResponse(login login.LoginDomainInterface) response.LoginResponse {
	return response.LoginResponse{
		ID:    login.GetID(),
		Name:  login.GetName(),
		Email: login.GetEmail(),
	}
}
