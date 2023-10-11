package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (login *loginHandler) LoginHandler(context *gin.Context) {
	context.JSON(http.StatusOK, "Login")

}
