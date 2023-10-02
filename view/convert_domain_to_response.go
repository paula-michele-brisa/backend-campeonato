package view

import (
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/response"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

// ConverteToDomainResponse covnerte o objeto domain para response
func ConverteToDomainResponse(userDomain user.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}
