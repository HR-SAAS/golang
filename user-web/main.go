package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
	"hr-saas-go/user-web/initial"
	"hr-saas-go/user-web/request"
)

func main() {
	app := gin.Default()
	initial.InitLog()
	initial.InitConfig()
	initial.InitApiRoute(app)
	// 初始化表单 验证
	initial.InitValidator()
	initial.InitRegister()
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

	host := fmt.Sprintf("%s:%d", global.Config.Host, global.Config.Port)
	zap.S().Debugf("server run in %s ", host)
	err := app.Run(host)
	if err != nil {
		zap.S().Panicf("error %s", err.Error())
	}
}
