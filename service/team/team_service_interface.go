package team

import "github.com/gin-gonic/gin"

type TeamDomainServiceInterface interface {
	CreateTeamService(context *gin.Context)
	DeleteTeamService(context *gin.Context)
	FindTeamByIDService(context *gin.Context)
	FindTeamsService(context *gin.Context)
	FindTotalTeamsService(context *gin.Context)
	UpdateTeamService(context *gin.Context)
}
