package login

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	login2 "github.com/paula-michele-brisa/backend-campeonato/domain/login"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/view"
	"net/http"
)

func (login *loginHandler) LoginHandler(context *gin.Context) {

	var loginRequest request.LoginRequest

	if err := context.BindJSON(loginRequest); err != nil {
		logger.Error("Erro ao efetuar login", err)
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	loginDomain := login2.NewLoginDomain(
		loginRequest.Email,
		loginRequest.Name,
		loginRequest.Password,
	)

	loginUser, err := login.loginService.LoginService(loginDomain)

	if err != nil {
		logger.Error("Erro trying to call LoginUser services", err)
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.ConverteToLoginDomainResponse(loginUser))

}
