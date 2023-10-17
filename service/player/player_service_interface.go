package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
)

type PlayerServiceInterface interface {
	CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr)
	FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr)
	FindTotalPlayer() (int32, *rest_err.RestErr)
	FindPlayerByID(id string) (int32, *rest_err.RestErr)
	UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr
	DeletePlayer(id string) *rest_err.RestErr
}
