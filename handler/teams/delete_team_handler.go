package teams

import "github.com/gin-gonic/gin"

// DeleteTeamHandler remove o time
func (team *teamHandler) DeleteTeamHandler(context *gin.Context) {
	context.JSON(204, gin.H{
		"message": "Time excluído",
	})
}
