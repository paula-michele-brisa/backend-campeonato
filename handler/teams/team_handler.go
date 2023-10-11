package teams

import (
	"github.com/paula-michele-brisa/backend-campeonato/service/team"
)

type teamHandler struct {
	teamService team.TeamDomainServiceInterface
}
