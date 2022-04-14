package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/company-web/api"
	"hr-saas-go/company-web/api/company"
	"hr-saas-go/company-web/api/department"
	"hr-saas-go/company-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	companyRouter := group.Group("/company").Use(middleware.JWTAuth())
	{
		companyRouter.GET("/", company.List)
		companyRouter.GET(":id", company.Show)
		companyRouter.POST("/", company.Create)
		companyRouter.PUT("/", company.Update)
		companyRouter.DELETE(":id", company.Delete)
	}
	departmentRouter := group.Group("/department").Use(middleware.JWTAuth())
	{
		departmentRouter.GET("/", department.List)
		departmentRouter.GET(":id", department.Show)
		departmentRouter.POST("/", department.Create)
		departmentRouter.PUT("/", department.Update)
		departmentRouter.DELETE(":id", department.Delete)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
