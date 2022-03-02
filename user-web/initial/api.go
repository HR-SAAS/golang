package initial

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/user-web/router"
)

func InitApiRoute(app *gin.Engine) {
	api := app.Group("v1")
	router.InitUserRouter(api)

}
