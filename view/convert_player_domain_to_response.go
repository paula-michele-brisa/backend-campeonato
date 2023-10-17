package view

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/request"
)

func ConvertPlayerDomainToResponse(playerDomain player.PlayerDomainInterface) request.PlayerRequest {
	return request.PlayerRequest{
		Weight:   playerDomain.GetWeight(),
		Name:     playerDomain.GetName(),
		Height:   playerDomain.GetHeight(),
		Number:   playerDomain.GetNumber(),
		Photo:    playerDomain.GetPhoto(),
		Position: playerDomain.GetPosition(),
		Age:      playerDomain.GetAge(),
		TeamID:   playerDomain.GetTeamID(),
	}
}
