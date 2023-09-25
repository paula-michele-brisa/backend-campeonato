package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/models/request"
	"net/http"
)

// RegisterUserHandler é responsável por retorar os dados do usuário
func RegisterUserHandler(context *gin.Context) {

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, error=%s", err))

		context.JSON(http.StatusBadRequest, restErr)
		return

	}

	fmt.Println(userRequest)

}
