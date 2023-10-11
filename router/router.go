package router

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/teams"
	"github.com/paula-michele-brisa/backend-campeonato/handler/user"
	"log"
)

// SetupRouter configura o servidor
func SetupRouter(userHandler user.UserHandlerInterface, teamHandler teams.TeamHandlerInterface) {
	router := gin.Default()

	InitializeRouters(router, userHandler, teamHandler)

	if err := router.Run(); err != nil {
		log.Fatal("Ocorreu um erro ao iniciar o servidor.")
	}

}
