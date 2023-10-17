package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
)

func NewPlayerRepository() PlayerRepositoryInterface {
	return &playerRepository{}
}

func (player *playerRepository) CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (player *playerRepository) FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (player *playerRepository) FindTotalPlayer() (int32, *rest_err.RestErr) { return 0, nil }

func (player *playerRepository) FindPlayerByID(id string) (player.PlayerDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (player *playerRepository) UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr {
	return nil
}

func (player *playerRepository) DeletePlayer(id string) *rest_err.RestErr {
	return nil
}
