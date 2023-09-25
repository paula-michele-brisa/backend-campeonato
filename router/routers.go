package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/router/games"
	"github.com/paula-michele-brisa/backend-campeonato/router/login"
	"github.com/paula-michele-brisa/backend-campeonato/router/player"
	"github.com/paula-michele-brisa/backend-campeonato/router/team"
	"github.com/paula-michele-brisa/backend-campeonato/router/user"
)

func routers(router *gin.Engine) {

	basePath := "api/v1"

	v1 := router.Group(basePath)
	{

		// Obter usuário
		v1.GET("/user", handler.GetUserHandler)
		user.UserRouter(v1)
		login.LoginRouter(v1)
		team.TeamRouter(v1)
		player.PlayerRouter(v1)
		games.GamesRouter(v1)

	}

}
