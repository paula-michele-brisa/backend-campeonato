package login

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
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

	loginUser, err := login.loginService.LoginService(loginRequest)

	context.JSON(http.StatusOK, "Login")

}
