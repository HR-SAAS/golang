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
		PostRouter.GET("/", post.List)
		PostRouter.GET("/:id", post.Show)
		PostRouter.POST("/", post.Create)
		PostRouter.PUT("/:id", post.Update)
		PostRouter.DELETE("/:id", post.Delete)
	}

	UserPostRouter := group.Group("/recruit/user_post").Use(middleware.JWTAuth())
	{
		UserPostRouter.GET("/", user_post.List)
		PostRouter.GET("/:id", post.Show)
		PostRouter.POST("/", post.Create)
		PostRouter.PUT("/:id", post.Update)
		PostRouter.DELETE("/:id", post.Delete)
	}
}

func InitHealth(group *gin.RouterGroup) {
	userGroup := group.Group("health")
	{
		userGroup.GET("/", api.Health)
	}
}
