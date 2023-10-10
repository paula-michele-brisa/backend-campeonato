package teams

import (
	"github.com/gin-gonic/gin"
)

// CreateTeamHandler cadastra um novo time
func (team *teamHandlerInterface) CreateTeamHandler(context *gin.Context) {

	context.JSON(201, gin.H{
		"message": "Time Cadastrado",
	})
}
