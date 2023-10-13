package login

func NewLoginDomain(email, password, name string) LoginDomainInterface {
	return &loginDomain{
		Email:    email,
		Name:     name,
		Password: password,
	}

}
