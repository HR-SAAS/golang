package initial

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/user-web/router"
)

func InitApiRoute(app *gin.Engine) {
	api := app.Group("api")
	r := api.Group("v1")

	router.InitUserRouter(r)
	router.InitCaptcha(r)
	router.InitHealth(api)
}
