package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	user2 "github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"github.com/paula-michele-brisa/backend-campeonato/router/team"
	"github.com/paula-michele-brisa/backend-campeonato/router/user"
)

// InitializeRouters inicializa as rotas
func InitializeRouters(router *gin.Engine, userHandler user2.UserHandlerInterface, teamHandler teams.TeamHandlerInterface) {

	v1 := router.Group("api/v1")

	user.UserRouters(v1, userHandler)
	team.TeamRouter(v1, teamHandler)

}
