package teams

import "github.com/gin-gonic/gin"

// GetTotalRegisteredTeamsHandler obtem a quantidade total de times cadastrados
func GetTotalRegisteredTeamsHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"total": "10",
	})
}
