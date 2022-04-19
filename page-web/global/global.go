package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/page-web/config"
	"hr-saas-go/page-web/proto"
)

var (
	Config            *config.ServerConfig = &config.ServerConfig{}
	NacosConfig       *config.NacosConfig  = &config.NacosConfig{}
	Trans             ut.Translator
	CompanyServCon    proto.CompanyClient
	DepartmentServCon proto.DepartmentClient
	UserServCon       proto.UserClient
	Debug             bool
)
