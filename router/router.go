package router

import "github.com/gin-gonic/gin"

// SetupRouter inicializa e configura as rotas e o servidor
func SetupRouter() {
	router := gin.Default()

	v1 := router.Group("api/v1")

	initializeRouters(v1)

	router.Run()

}
