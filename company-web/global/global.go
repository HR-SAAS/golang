package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/company-web/config"
	"hr-saas-go/company-web/proto"
)

var (
	Config            *config.ServerConfig = &config.ServerConfig{}
	NacosConfig       *config.NacosConfig  = &config.NacosConfig{}
	Trans             ut.Translator
	CompanyServCon    proto.CompanyClient
	DepartmentServCon proto.DepartmentClient
	Debug             bool
)
