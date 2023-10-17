package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}

func FindPlayers(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}

func FindTotalPlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")

}

func FindPlayerByID(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")

}

func UpdatePlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}

func DeletePlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}
