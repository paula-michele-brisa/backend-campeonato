package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	user2 "github.com/paula-michele-brisa/backend-campeonato/entity/user"
)

func ConvertDomainToEntity(userDomain user.UserDomainInterface) *user2.UserEntity {
	return &user2.UserEntity{
		ID:       userDomain.GetID(),
		Name:     userDomain.GetName(),
		Password: userDomain.GetPassword(),
		Email:    userDomain.GetEmail(),
	}

}
