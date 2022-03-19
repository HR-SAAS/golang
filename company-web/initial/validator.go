package initial

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translate "github.com/go-playground/validator/v10/translations/en"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"hr-saas-go/company-web/global"
	"reflect"
	"strings"
)

func initTrans(local string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("json")
			if name == "-" {
				return ""
			}
			return strings.ToLower(name)
		})

		// 国际化
		enT := en.New()
		zhT := zh.New()

		u := ut.New(enT, zhT, enT)

		global.Trans, ok = u.GetTranslator(local)
		if ok {
			switch local {
			case "en":
				return en_translate.RegisterDefaultTranslations(v, global.Trans)
			case "zh":
				return zh_translate.RegisterDefaultTranslations(v, global.Trans)
			default:
				return zh_translate.RegisterDefaultTranslations(v, global.Trans)
			}
		}
	}
	return errors.Errorf("国际化初始化失败 %s", local)
}

func InitValidator() {
	local := global.Config.Local
	if local == "" {
		local = "zh"
	}
	zap.S().Debugf("初始化表单验证 %s", local)
	err := initTrans(local)
	if err != nil {
		panic(err)
	}
}
