package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/championship"
	"github.com/paula-michele-brisa/backend-campeonato/router/login"
	"github.com/paula-michele-brisa/backend-campeonato/router/player"
	"github.com/paula-michele-brisa/backend-campeonato/router/team"
	"github.com/paula-michele-brisa/backend-campeonato/router/user"
)

func routers(router *gin.Engine) {

	basePath := "api/v1"

	v1 := router.Group(basePath)
	{
		user.UserRouter(v1)
		login.LoginRouter(v1)
		team.TeamRouter(v1)
		player.PlayerRouter(v1)

		// Obtem os jogos do campeonato
		v1.GET("/games", championship.GetGamesHandler)

		// Gera os jogos
		v1.POST("/championship", championship.GenerateChampionship)

	}

}
