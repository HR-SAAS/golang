package initial

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/router"
)

func InitApiRoute(app *gin.Engine) {
	api := app.Group("api")
	r := api.Group("v1")

	router.InitRouter(r)
	router.InitHealth(api)
}
