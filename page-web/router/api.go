package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/page-web/api"
	"hr-saas-go/page-web/api/company"
	"hr-saas-go/page-web/api/department"
	"hr-saas-go/page-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	companyRouter := group.Group("/company").Use(middleware.JWTAuth())
	{
		companyRouter.GET("/", company.List)
		companyRouter.GET("/my", company.MyCompany)
		companyRouter.GET("/:id", company.Show)
		companyRouter.GET("/:id/users", company.GetCompanyUsers)
		companyRouter.POST("/", company.Create)
		companyRouter.PUT("/:id", company.Update)
		companyRouter.DELETE("/:id", company.Delete)
		// 用户相关
	}
	departmentRouter := group.Group("/department").Use(middleware.JWTAuth())
	{
		departmentRouter.GET("/", department.List)
		departmentRouter.GET("/my", department.MyDepartment)
		departmentRouter.GET("/:id", department.Show)
		departmentRouter.POST("/", department.Create)
		departmentRouter.PUT("/:id", department.Update)
		departmentRouter.DELETE("/:id", department.Delete)
		// 统计相关
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
