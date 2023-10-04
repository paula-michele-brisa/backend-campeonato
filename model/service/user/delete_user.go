package user

import "github.com/paula-michele-brisa/backend-campeonato/config/rest_err"

func (user *userDomainService) DeleteUser(userId string) *rest_err.RestErr {

	err := user.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
