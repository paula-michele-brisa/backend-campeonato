package login

func NewLoginDomain(name, email, token string) LoginDomainInterface {
	return &loginDomain{
		Email: email,
		Name:  name,
		Token: token,
	}

}
