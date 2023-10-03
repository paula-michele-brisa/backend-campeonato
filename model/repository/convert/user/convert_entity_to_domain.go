package user

import (
	user2 "github.com/paula-michele-brisa/backend-campeonato/model/repository/entity/user"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func ConvertEntityToDomain(userEntity user2.UserEntity) user.UserDomainInterface {
	domain := user.NewUserDomain(
		userEntity.Email, userEntity.Password, userEntity.Name)

	domain.SetID(userEntity.ID)
	return domain
}
