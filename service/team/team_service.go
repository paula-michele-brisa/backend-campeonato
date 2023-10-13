package team

import (
	"github.com/paula-michele-brisa/backend-campeonato/repository/team"
)

type teamService struct {
	repository team.TeamRepositoryInterface
}
