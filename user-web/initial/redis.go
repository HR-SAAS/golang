package initial

import (
	"context"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/utils"
)

func InitRedis() {
	global.Rdb = utils.NewRedis()
	_, err := global.Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
