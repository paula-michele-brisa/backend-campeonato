package router

import "github.com/gin-gonic/gin"

// SetupRouter inicializa o servidor
func SetupRouter() {
	router := gin.Default() // Criando instancia do Gin Router e utilizando as configurações padrão

	routers(router)

	router.Run()

}
