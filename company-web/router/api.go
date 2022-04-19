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
		//获取列表
		companyRouter.GET("/", company.List)
		// 我加入的公司
		companyRouter.GET("/my", company.MyCompany)
		// 统计
		companyRouter.GET("/count", company.CompanyCount)
		//详情
		companyRouter.GET("/:id", company.Show)

		// 获取用户列表
		companyRouter.GET("/:id/users", company.GetCompanyUsers)
		// 创建
		companyRouter.POST("/", company.Create)
		// 修改
		companyRouter.PUT("/:id", company.Update)
		// 删除
		companyRouter.DELETE("/:id", company.Delete)
		// 用户相关
	}
	departmentRouter := group.Group("/department").Use(middleware.JWTAuth())
	{
		departmentRouter.GET("/", department.List)
		departmentRouter.GET("/my", department.MyDepartment)
		departmentRouter.GET("/count", company.MyCompany)
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
