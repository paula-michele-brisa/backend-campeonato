package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/model/service/user"
)

func NewUserHandlerInterface(serviceInterface user.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		service: serviceInterface,
	}
}

type UserHandlerInterface interface {
	CreateUserHandler(context *gin.Context)
	DeleteUserHandler(context *gin.Context)
	FindUserByIDHandler(context *gin.Context)
	FindUserByEmailHandler(context *gin.Context)
	UpdateUserHandler(context *gin.Context)
}

type userHandlerInterface struct {
	service user.UserDomainService
}
