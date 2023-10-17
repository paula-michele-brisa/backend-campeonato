package player

import (
	"github.com/gin-gonic/gin"
	player2 "github.com/paula-michele-brisa/backend-campeonato/domain/player"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
	"github.com/paula-michele-brisa/backend-campeonato/service/player"
	"github.com/paula-michele-brisa/backend-campeonato/view"
	"net/http"
)

func NewPlayerHandler(service player.PlayerServiceInterface) PlayerHandlerInterface {
	return &playerHandler{playerService: service}
}

func (player *playerHandler) CreatePlayer(context *gin.Context) {
	var playerRequest request.PlayerRequest

	if err := context.ShouldBindJSON(&playerRequest); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao cadastrar jogador. Erro interno!")
		return
	}

	playerDomain := player2.NewPlayerDomain(
		playerRequest.Name,
		playerRequest.Photo,
		playerRequest.Position,
		playerRequest.TeamID,
		playerRequest.Number,
		playerRequest.Weight,
		playerRequest.Age,
		playerRequest.Height,
	)

	result, err := player.playerService.CreatePlayer(playerDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao cadastrar jogador. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, view.ConvertPlayerDomainToResponse(result))
}

func (player *playerHandler) FindPlayers(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}

func (player *playerHandler) FindTotalPlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")

}

func (player *playerHandler) FindPlayerByID(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")

}

func (player *playerHandler) UpdatePlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}

func (player *playerHandler) DeletePlayer(context *gin.Context) {
	context.JSON(http.StatusOK, "Jogadores ok!")
}
