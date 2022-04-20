package initial

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/message-web/router"
)

func InitApiRoute(app *gin.Engine) {
	api := app.Group("api")
	r := api.Group("v1")

	router.InitCompanyRouter(r)
	router.InitHealth(api)
}
