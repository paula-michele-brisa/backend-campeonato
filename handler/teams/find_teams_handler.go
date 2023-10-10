package teams

import "github.com/gin-gonic/gin"

// GetTeamsHandler retorna a lista de times cadastrados
func (team *teamHandlerInterface) FindTeams(context *gin.Context) {
	context.JSON(200, gin.H{
		"time": "GPS futebol clube",
	})

}
