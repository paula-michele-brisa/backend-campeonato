package view

import (
	"github.com/paula-michele-brisa/backend-campeonato/domain/team"
	"github.com/paula-michele-brisa/backend-campeonato/handler/models/response"
)

func ConvertTeamDomainToResponse(team team.TeamDomainInterface) response.TeamResponse {
	return response.TeamResponse{
		Coach:      team.GetCoach(),
		Name:       team.GetName(),
		BadgePhoto: team.GetBadgePhoto(),
		City:       team.GetCity(),
		Website:    team.GetWebSite(),
		ID:         team.GetID(),
	}
}
