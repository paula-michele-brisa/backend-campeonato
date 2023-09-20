package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default() // Criando instancia do Gin Router e utilizando as configurações padrão

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})

	})

	return router

}
