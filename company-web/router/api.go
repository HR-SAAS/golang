package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/company-web/api"
	company "hr-saas-go/company-web/api/company"
	"hr-saas-go/company-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	userGroup := group.Group("companies").Use(middleware.JWTAuth())
	{
		userGroup.GET("/", company.List)
		userGroup.GET(":id", company.Show)
		userGroup.POST("/", company.Create)
		userGroup.PUT("/", company.Update)
		userGroup.DELETE(":id", company.Delete)
	}

}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
