package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/user-web/api"
	"hr-saas-go/user-web/middleware"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("users").Use(middleware.JWTAuth())
	{
		userGroup.GET("/", api.GetUserList)
		userGroup.GET(":id", api.GetUser)
	}
	authGroup := group.Group("auth")
	{
		authGroup.POST("loginByMobile", api.LoginByMobile)
	}
}

func InitCaptcha(group *gin.RouterGroup) {
	userGroup := group.Group("captcha")
	{
		userGroup.GET("/img", api.GetCaptcha)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
