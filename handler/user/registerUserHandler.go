package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/validation"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
)

// RegisterUserHandler é responsável por retorar os dados do usuário
func RegisterUserHandler(context *gin.Context) {

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		context.JSON(restErr.Code, restErr)
		return

	}

	fmt.Println(userRequest)

}
