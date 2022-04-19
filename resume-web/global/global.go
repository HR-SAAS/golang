package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/resume-web/config"
	"hr-saas-go/resume-web/proto"
)

var (
	Config             *config.ServerConfig = &config.ServerConfig{}
	NacosConfig        *config.NacosConfig  = &config.NacosConfig{}
	Trans              ut.Translator
	ResumeServCon      proto.ResumeClient
	ResumeCountServCon proto.ResumeCounterServiceClient
	Debug              bool
)
