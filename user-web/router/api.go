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
		userGroup.GET("info", api.GetInfo)
		userGroup.GET(":id", api.Show)
	}
	authGroup := group.Group("auth")
	{
		authGroup.POST("loginByPassword", api.LoginByPassword)
		authGroup.POST("loginByCode", api.LoginByCode)
		authGroup.POST("register", api.Register)
	}
}

func InitCaptcha(group *gin.RouterGroup) {
	userGroup := group.Group("captcha")
	{
		userGroup.GET("/img", api.GetCaptcha)
		userGroup.POST("/sms", api.GetSmsCaptcha)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
