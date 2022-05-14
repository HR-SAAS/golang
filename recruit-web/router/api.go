package router

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/api"
	"hr-saas-go/recruit-web/api/post"
	"hr-saas-go/recruit-web/api/user_post"
	"hr-saas-go/recruit-web/middleware"
)

func InitRouter(group *gin.RouterGroup) {
	PostRouter := group.Group("/recruit/post").Use(middleware.JWTAuth())
	{
		// 获取岗位列表
		PostRouter.GET("/", post.List)
		// 获取岗位详情
		PostRouter.GET("/:id", post.Show)
		// 创建岗位
		PostRouter.POST("/", post.Create)
		//更新岗位
		PostRouter.PUT("/:id", post.Update)
		// 删除岗位
		PostRouter.DELETE("/:id", post.Delete)
	}

	UserPostRouter := group.Group("/recruit/user_post").Use(middleware.JWTAuth())
	{
		// 获取投递列表
		UserPostRouter.GET("/", user_post.List(""))
		// 获取我的投递记录
		UserPostRouter.GET("/me", user_post.List("user"))
		// 获取岗位的投递记录
		UserPostRouter.GET("/post", user_post.List("post"))
		// 获取企业的投递记录
		UserPostRouter.GET("/company", user_post.List("company"))

		// 检测
		UserPostRouter.GET("/check", user_post.CheckCanPost)
		// 获取详情
		UserPostRouter.GET("/:id", user_post.Show)
		// 投递
		UserPostRouter.POST("/", user_post.Create)
		// 修改状态
		UserPostRouter.PUT("/:id", user_post.Update)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
