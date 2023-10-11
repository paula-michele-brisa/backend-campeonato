package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/service/user"
)

type userHandlerInterface struct {
	service user.UserDomainService
}
