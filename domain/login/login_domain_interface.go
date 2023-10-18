package login

import "github.com/paula-michele-brisa/backend-campeonato/config/rest_err"

type LoginDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	SetPassword(password string)
	GetPassword() string
	SetID(id string)
	GenerateToken() (string, *rest_err.RestErr)
}
