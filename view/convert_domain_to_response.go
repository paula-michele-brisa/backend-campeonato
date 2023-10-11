package view

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/user"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/response"
)

// ConverteToDomainResponse covnerte o objeto domain para response
func ConverteToDomainResponse(userDomain user.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}
