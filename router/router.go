package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter inicializa e configura as rotas e o servidor
func SetupRouter() {
	router := gin.Default()

	v1 := router.Group("api/v1")

	initializeRouters(v1)

	if err := router.Run(); err != nil {
		log.Fatal(err)

	}

}
