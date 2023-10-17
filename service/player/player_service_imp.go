package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
)

func CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr) { return nil, nil }

func FindTotalPlayer() (int32, *rest_err.RestErr) { return 0, nil }

func FindPlayerByID(id string) (int32, *rest_err.RestErr) {
	return 0, nil
}

func UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr {
	return nil
}

func DeletePlayer(id string) *rest_err.RestErr {
	return nil
}
