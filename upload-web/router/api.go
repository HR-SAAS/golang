package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/upload-web/api"
	"hr-saas-go/upload-web/middleware"
)

func InitCompanyRouter(group *gin.RouterGroup) {
	group.Use(middleware.ReadFromJwt())
	{
		group.POST("/upload/images", api.UploadFile("images"))
		group.POST("/upload/resumes", api.UploadFile("resumes"))
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
