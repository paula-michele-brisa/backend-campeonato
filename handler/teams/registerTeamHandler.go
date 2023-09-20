package teams

import "github.com/gin-gonic/gin"

// RegisterTeamHandler cadastra um novo time
func RegisterTeamHandler(context *gin.Context) {
	context.JSON(201, gin.H{
		"message": "Time Cadastrado",
	})
}
