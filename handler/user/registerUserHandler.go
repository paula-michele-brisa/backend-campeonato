package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"net/http"
)

var (
	UserDomainInterface user.UserDomainInterface
)

// RegisterUserHandler é responsável por retorar os dados do usuário
func RegisterUserHandler(context *gin.Context) {

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := user.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name)

	if err := domain.CreateUser(); err != nil {
		context.JSON(err.Code, err)
	}

	context.String(http.StatusOK, "")

}
