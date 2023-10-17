package player

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/player"
	player2 "github.com/paula-michele-brisa/backend-campeonato/entity/player"
)

func ConverPlayerEntityToDomain(playerEntity player2.PlayerEntity) player.PlayerDomainInterface {

	domain := player.NewPlayerDomain(
		playerEntity.Name,
		playerEntity.TeamID,
		playerEntity.Photo,
		playerEntity.Position,
		playerEntity.Weight,
		playerEntity.Age,
		playerEntity.Number,
		playerEntity.Height,
	)

	domain.SetID(playerEntity.Id)

	return domain
}
