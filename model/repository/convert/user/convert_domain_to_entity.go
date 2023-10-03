package user

import (
	user2 "github.com/paula-michele-brisa/backend-campeonato/model/repository/entity/user"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func ConvertDomainToEntity(userDomain user.UserDomainInterface) *user2.UserEntity {
	return &user2.UserEntity{
		ID:       userDomain.GetID(),
		Name:     userDomain.GetName(),
		Password: userDomain.GetPassword(),
		Email:    userDomain.GetEmail(),
	}

}
