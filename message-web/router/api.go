package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/message-web/api"

	"hr-saas-go/message-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	companyRouter := group.Group("/message").Use(middleware.JWTAuth())
	{
		companyRouter.GET("/", api.List)
		companyRouter.GET("/count", api.Count)
		companyRouter.GET("/:id", api.Show)
		companyRouter.POST("/send", api.SendTo)
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
