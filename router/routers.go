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

		// Obter o total de times cadastrados
		v1.GET("/team_count", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "10",
			})
		})

		// Listar os times cadastrados
		v1.GET("/teams", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"time": "GPS futebol clube",
			})
		})

		// Cadastra um novo time
		v1.POST("/team", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Time Cadastrado",
			})
		})

		// Editar/Atualizar time
		v1.PUT("/team/:id", func(context *gin.Context) {
			context.JSON(400, gin.H{
				"message": "Time atualizado",
			})
		})

		// Deleta um time
		v1.DELETE("/time/:id", func(context *gin.Context) {
			context.JSON(204, gin.H{
				"message": "Usuário logado",
			})
		})

		// Obtem um time pelo id
		v1.GET("/team/:id", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "10",
			})
		})

	}

}
