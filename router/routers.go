package router

import "github.com/gin-gonic/gin"

func routers(router *gin.Engine) {

	basePath := "api/v1"

	v1 := router.Group(basePath)
	{
		// Obter usuário
		v1.GET("/user", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"usario": "xx1254",
			})
		})

		// Registrar usuário
		v1.POST("/user", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Usuário cadastrado",
			})
		})

		// Efetuar login
		v1.POST("/login", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Usuário logado",
			})
		})

	}

}
