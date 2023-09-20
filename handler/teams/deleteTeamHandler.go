package teams

import "github.com/gin-gonic/gin"

// DeleteTeamHandler remove o time
func DeleteTeamHandler(context *gin.Context) {
	context.JSON(204, gin.H{
		"message": "Time excluído",
	})
}
