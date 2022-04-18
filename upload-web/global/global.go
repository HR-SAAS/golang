package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/upload-web/config"
)

var (
	Config      *config.ServerConfig = &config.ServerConfig{}
	NacosConfig *config.NacosConfig  = &config.NacosConfig{}
	Trans       ut.Translator
	Debug       bool
)
