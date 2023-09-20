package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // Criando instancia do Gin Router e utilizando as configurações padrão

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})

	})

	router.Run() // listen and server on 0.0.0:8080

}
