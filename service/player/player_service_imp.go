package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
)

func NewPlayerService() PlayerServiceInterface {
	return &playerService{}
}

func (player *playerService) CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr) {
	return player.playerRepository.CreatePlayer(playerDomain)
}

func (player *playerService) FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr) {
	return player.playerRepository.FindPlayers()
}

func (player *playerService) FindTotalPlayer() (int32, *rest_err.RestErr) {
	return player.playerRepository.FindTotalPlayer()
}

func (player *playerService) FindPlayerByID(id string) (player.PlayerDomainInterface, *rest_err.RestErr) {
	return player.playerRepository.FindPlayerByID(id)
}

func (player *playerService) UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr {
	return player.playerRepository.UpdatePlayer(id, playerDomain)
}

func (player *playerService) DeletePlayer(id string) *rest_err.RestErr {
	return player.playerRepository.DeletePlayer(id)
}
