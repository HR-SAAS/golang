package initial

import (
	"github.com/gin-gonic/gin"
	"hr-saas-go/upload-web/router"
)

func InitApiRoute(app *gin.Engine) {
	api := app.Group("api")
	r := api.Group("v1")

	router.InitCompanyRouter(r)
	router.InitHealth(api)
}
