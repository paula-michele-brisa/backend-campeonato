package user

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
	"github.com/paula-michele-brisa/backend-campeonato/view"
	"go.uber.org/zap"
	"net/http"
)

// UpdateUserHandler retorna os dados do usuário
func (userHandler *userHandlerInterface) UpdateUserHandler(context *gin.Context) {
	logger.Info("Init UpdateUserHandler handler", zap.String("journey", "UpdateUserHandler"))

	userID := context.Param("userId")
	var userRequest request.UserUpdateRequest

	userUpdateDomain := user.NewUpdateUserDomain(
		userRequest.Name,
		userRequest.Email)

	err := userHandler.service.UpdateUser(userID, userUpdateDomain)
	if err != nil {
		logger.Error("Erro trying to call UpdateUserHandler services", err)
		context.JSON(err.Code, err)
		return

	}

	logger.Info("UpdateUserHandler handler excuted successfully", zap.String("journey", "findUserByID"))

	context.JSON(http.StatusOK, view.ConverteToDomainResponse(userUpdateDomain))

}
