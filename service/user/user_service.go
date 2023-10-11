package user

import (
	user2 "github.com/paula-michele-brisa/backend-campeonato/repository/user"
)

type userDomainService struct {
	userRepository user2.UserRepository
}
