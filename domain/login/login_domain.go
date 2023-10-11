package login

type loginDomain struct {
	Name  string
	Email string
	Token string
}

func (login *loginDomain) GetName() string {
	return login.Name
}

func (login *loginDomain) GetEmail() string {
	return login.Email
}

func (login *loginDomain) GetToken() string {
	return login.Token
}
