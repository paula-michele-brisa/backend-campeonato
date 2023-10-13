package login

type loginDomain struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (login *loginDomain) GetID() string {
	return login.ID
}

func (login *loginDomain) GetName() string {
	return login.Name
}

func (login *loginDomain) GetEmail() string {
	return login.Email
}

func (user *loginDomain) SetPassword(password string) {
	user.Password = password
}
