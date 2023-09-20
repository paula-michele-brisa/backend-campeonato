package teams

import "github.com/gin-gonic/gin"

// GetTeamHandler obtém um time pelo id
func GetTeamHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"total": "10",
	})
}
