package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/user-web/config"
)

var (
	Config *config.ServerConfig = &config.ServerConfig{}
	Trans  ut.Translator
)
