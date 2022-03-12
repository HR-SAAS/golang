package initial

import (
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
)

func InitRegister() {
	config := global.Config
	err := utils.Register(config.Name, config.Name, config.Host, config.Port, []string{
		"user-web", "golang", "web",
	})
	if err != nil {
		panic(err)
	}
}
