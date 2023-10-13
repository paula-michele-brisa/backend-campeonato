package login

type LoginDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
}
