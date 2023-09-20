package user

import "github.com/gin-gonic/gin"

// GetUserHandler é responsável por retorar os dados do usuário
func GetUserHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"usario": "xx1254",
	})
}
