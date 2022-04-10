package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/initial"
	"hr-saas-go/user-web/request"
	"hr-saas-go/user-web/utils"
	"os"
	"os/signal"
	"syscall"
)

func generatePort() {

}

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

	// 自动生成host&port
	err := register.Register(config.Name, id, config.Host, config.Port, []string{
		"user-web", "golang", "web",
	})
	if err != nil {
		zap.S().Errorf("注册失败: %s", err)
	}

	initial.InitialCon()
	// 注册自定义表单验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", request.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 必须为正确的手机号格式", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

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
