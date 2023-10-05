package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"github.com/paula-michele-brisa/backend-campeonato/view"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface user.UserDomainInterface
)

// CreateUserHandler cadastra o usuário
func (userHandler *userHandlerInterface) CreateUserHandler(context *gin.Context) {

	// Objeto com os atributos necessários para fazer a criação do usuário
	var userRequest request.UserRequest

	logger.Info(userRequest.Name)

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	registeredUser, err := userHandler.service.FindUserByEmailServices(userRequest.Email)

	if registeredUser != nil {
		context.JSON(http.StatusBadRequest, "msg: Usuário já cadastrado no sistema!")
		return
	}

	domain := user.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name)

	domainResult, err := userHandler.service.CreateUser(domain)

	if err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(domainResult))
}
