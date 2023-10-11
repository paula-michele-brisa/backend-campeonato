package team

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
)

type TeamDomainServiceInterface interface {
	CreateTeamService(team TeamDomainServiceInterface) (TeamDomainServiceInterface, *rest_err.RestErr)
	DeleteTeamService(context *gin.Context)
	FindTeamByIDService(context *gin.Context)
	FindTeamsService(context *gin.Context)
	FindTotalTeamsService(context *gin.Context)
	UpdateTeamService(context *gin.Context)
}
