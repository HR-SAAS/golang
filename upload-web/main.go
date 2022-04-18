package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"hr-saas-go/upload-web/global"
	"hr-saas-go/upload-web/initial"
	"hr-saas-go/upload-web/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := gin.Default()
	initial.InitLog()
	initial.InitConfig()
	initial.InitApiRoute(app)
	// 初始化表单 验证
	initial.InitValidator()

	//注册
	config := global.Config
	register := utils.NewRegister(config.ConsulConfig.Host, config.ConsulConfig.Port)
	// 自动生成id
	u2 := uuid.NewV4()
	id := fmt.Sprintf("%v", u2)

	// 获取端口和ip地址
	err := register.Register(config.Name, id, config.Host, config.Port, []string{
		"upload-web", "golang", "web",
	})
	if err != nil {
		zap.S().Errorf("注册失败: %s", err)
	}

	initial.InitialCon()

	host := fmt.Sprintf("%s:%d", config.Host, config.Port)
	zap.S().Debugf("server run in %s ", host)

	go func() {
		if err = app.Run(host); err != nil {
			zap.S().Panicf("error %s", err.Error())
		}
	}()

	//signal
	sigC := make(chan os.Signal)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
	<-sigC
	//注销
	if err = register.Deregister(id); err != nil {
		zap.S().Errorf("注销失败 %s", err)
	} else {
		zap.S().Infof("注销成功")
	}
}
