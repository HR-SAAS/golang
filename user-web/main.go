package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hr-saas-go/user-web/initial"
)

func main() {
	port := 8082
	host := fmt.Sprintf(":%d", port)
	app := gin.Default()
	initial.InitLog()
	initial.InitConfig()
	initial.InitApiRoute(app)
	// 初始化表单 验证
	initial.InitValidator()
	zap.S().Debugf("server run in %s ", host)
	err := app.Run(host)
	if err != nil {
		zap.S().Panicf("error %s", err.Error())
	}
}
