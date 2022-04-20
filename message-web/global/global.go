package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/message-web/config"
	"hr-saas-go/message-web/proto"
)

var (
	Config                *config.ServerConfig = &config.ServerConfig{}
	NacosConfig           *config.NacosConfig  = &config.NacosConfig{}
	Trans                 ut.Translator
	UserMessageServCon    proto.UserMessageClient
	MessageCounterService proto.MessageCounterClient
	Debug                 bool
)
