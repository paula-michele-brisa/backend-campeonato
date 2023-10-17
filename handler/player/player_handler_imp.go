package player

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"net/http"
)

func CreatePlayer(context *gin.Context) {
	var playerRequest request.PlayerRequest

	if err := context.ShouldBindJSON(&playerRequest); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao cadastrar jogador. Erro interno!")
		return
	}

	//playerDomain := player.NewPlayerDomain(
	//	playerRequest.Name,
	//	playerRequest.Photo,
	//	playerRequest.Position,
	//	playerRequest.TeamID,
	//	playerRequest.Number,
	//	playerRequest.Weight,
	//	playerRequest.Age,
	//	playerRequest.Height,
	//)

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
