package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"net/http"
)

// DeleteUserHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) DeleteUserHandler(context *gin.Context) {
	userID := context.Param("userId")

	err := userHandler.service.DeleteUser(userID)

	if err != nil {
		if err != nil {
			logger.Error("Erro trying to call DeleteUserHandler services", err)
			context.JSON(err.Code, err)
			return
		}
	}

	context.JSON(http.StatusOK, "Usuário excluído com sucesso!")

}
