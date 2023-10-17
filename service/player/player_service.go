package player

import "github.com/paula-michele-brisa/backend-campeonato/repository/player"

type playerService struct {
	playerRepository player.PlayerRepositoryInterface
}
