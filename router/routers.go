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
				"message": "Time excluído",
			})
		})

		// Obtem um time pelo id
		v1.GET("/team/:id", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "10",
			})
		})

		// Obtem os jogadores cadastrados
		v1.GET("/payers", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "10",
			})
		})

		// Cadastra um novo jogador
		v1.POST("/player", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Player Cadastrado",
			})
		})

		// Editar/Atualizar jogador
		v1.PUT("/player/:id", func(context *gin.Context) {
			context.JSON(400, gin.H{
				"message": "Jogador atualizado",
			})
		})

		// Deleta um time
		v1.DELETE("/player/:id", func(context *gin.Context) {
			context.JSON(204, gin.H{
				"message": "Jogador excluído",
			})
		})

		// Obtem os jogos do campeonato
		v1.GET("/games", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"total": "GPS x 4 andar",
			})
		})

		// Gera os jogos
		v1.POST("/championship", func(context *gin.Context) {
			context.JSON(201, gin.H{
				"message": "Campeonato gerado com sucesso!",
			})
		})

	}

}
