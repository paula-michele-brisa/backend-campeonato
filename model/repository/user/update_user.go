package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func (user *userRespository) UpdateUser(id string, domainInterface user.UserDomainInterface) *rest_err.RestErr {
	return nil
}
