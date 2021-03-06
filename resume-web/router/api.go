package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/resume-web/api"

	"hr-saas-go/resume-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	companyRouter := group.Group("/resume").Use(middleware.JWTAuth())
	{
		companyRouter.GET("/", api.List)
		companyRouter.GET("/count", api.Count)
		companyRouter.GET("/:id", api.Show)
		companyRouter.POST("/", api.Create)
		companyRouter.PUT("/:id", api.Update)
		companyRouter.DELETE("/:id", api.Delete)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
