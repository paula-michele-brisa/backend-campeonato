package login

import "github.com/gin-gonic/gin"

func LoginRouter(router *gin.RouterGroup) {

	loginGroup := router.Group("/login")

	loginGroup.GET("/", nil)

}
