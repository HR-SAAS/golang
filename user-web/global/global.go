package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/user-web/config"
	"hr-saas-go/user-web/proto"
)

var (
	Config      = &config.ServerConfig{}
	NacosConfig = &config.NacosConfig{}
	Trans       ut.Translator
	UserServCon proto.UserClient
	Debug       bool
)
