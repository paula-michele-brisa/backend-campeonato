package user

import (
	login2 "github.com/paula-michele-brisa/backend-campeonato/domain/login"
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	"github.com/paula-michele-brisa/backend-campeonato/entity/login"
	user2 "github.com/paula-michele-brisa/backend-campeonato/entity/user"
)

func ConvertEntityToDomain(userEntity user2.UserEntity) user.UserDomainInterface {
	domain := user.NewUserDomain(
		userEntity.Email, userEntity.Password, userEntity.Name)

	domain.SetID(userEntity.ID)
	return domain
}

func ConvertLoginEntityToDomain(loginEntity login.LoginEntity) login2.LoginDomainInterface {
	domain := login2.NewLoginDomain(
		loginEntity.Email, loginEntity.Name, loginEntity.Password)

	domain.SetID(loginEntity.ID)
	return domain
}
