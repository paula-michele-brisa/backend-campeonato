package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface user.UserDomainInterface
)

// CreateUserHandler é responsável por retorar os dados do usuário
func CreateUserHandler(context *gin.Context) {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := user.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name)

	if err := domain.CreateUser(); err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	context.String(http.StatusOK, "")

}
