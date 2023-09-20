package teams

import "github.com/gin-gonic/gin"

func DeleteTeamHandler(context *gin.Context) {
	context.JSON(204, gin.H{
		"message": "Time excluído",
	})
}
