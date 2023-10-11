package team

import "github.com/paula-michele-brisa/backend-campeonato/repository/team"

type teamDomainService struct {
	teamRespository team.TeamRepositoryInterface
}
