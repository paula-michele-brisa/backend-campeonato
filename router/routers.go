package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/router/games"
	"github.com/paula-michele-brisa/backend-campeonato/router/login"
	"github.com/paula-michele-brisa/backend-campeonato/router/player"
	"github.com/paula-michele-brisa/backend-campeonato/router/team"
	"github.com/paula-michele-brisa/backend-campeonato/router/user"
)

func initializeRouters(router *gin.RouterGroup) {

	user.UserRouter(router)
	login.LoginRouter(router)
	team.TeamRouter(router)
	player.PlayerRouter(router)
	games.GamesRouter(router)

}
