package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/login"
	player2 "github.com/paula-michele-brisa/backend-campeonato/handler/player"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	login2 "github.com/paula-michele-brisa/backend-campeonato/router/login"
	"github.com/paula-michele-brisa/backend-campeonato/router/player"
	"github.com/paula-michele-brisa/backend-campeonato/router/team"
	"github.com/paula-michele-brisa/backend-campeonato/router/user"
)

// InitializeRouters inicializa as rotas
func InitializeRouters(router *gin.Engine, userHandler user2.UserHandlerInterface, teamHandler teams.TeamHandlerInterface, loginHandler login.LoginHandlerInterface, playerHandler player2.PlayerHandlerInterface) {

	v1 := router.Group("api/v1")

	login2.LoginRouter(v1, loginHandler)
	user.UserRouters(v1, userHandler)
	team.TeamRouter(v1, teamHandler)
	player.PLayerRouter(v1, playerHandler)

}
