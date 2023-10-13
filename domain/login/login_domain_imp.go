package login

func NewLoginDomain(email, name string) LoginDomainInterface {
	return &loginDomain{
		Email: email,
		Name:  name,
	}

}
