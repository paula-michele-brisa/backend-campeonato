package teams

import "github.com/gin-gonic/gin"

// UpdateTeamHandler atualiza o time pelo id
func (team *teamHandlerInterface) UpdateTeamHandler(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "Time atualizado",
	})

}
