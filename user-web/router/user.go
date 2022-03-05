package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/user-web/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("users")
	{
		userGroup.GET("/", api.GetUserList)
		userGroup.GET(":id", api.GetUser)
	}
	authGroup := group.Group("auth")
	{
		authGroup.POST("loginByMobile", api.LoginByMobile)
	}
}
