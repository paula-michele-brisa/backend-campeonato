package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
	player2 "github.com/paula-michele-brisa/backend-campeonato/entity/player"
)

func ConvertPlayerDomainToEntity(playerDomain player.PlayerDomainInterface) *player2.PlayerEntity {
	return &player2.PlayerEntity{
		Name:     playerDomain.GetName(),
		Photo:    playerDomain.GetPhoto(),
		Height:   playerDomain.GetHeight(),
		Weight:   playerDomain.GetWeight(),
		Age:      playerDomain.GetAge(),
		Position: playerDomain.GetPosition(),
		Number:   playerDomain.GetNumber(),
		TeamID:   playerDomain.GetTeamID(),
	}
}
