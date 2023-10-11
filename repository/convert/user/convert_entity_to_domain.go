package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	user2 "github.com/paula-michele-brisa/backend-campeonato/entity/user"
)

func ConvertEntityToDomain(userEntity user2.UserEntity) user.UserDomainInterface {
	domain := user.NewUserDomain(
		userEntity.Email, userEntity.Password, userEntity.Name)

	domain.SetID(userEntity.ID)
	return domain
}
