package global

import (
	ut "github.com/go-playground/universal-translator"
	"hr-saas-go/recruit-web/config"
	"hr-saas-go/recruit-web/proto"
)

var (
	Config                       *config.ServerConfig = &config.ServerConfig{}
	NacosConfig                  *config.NacosConfig  = &config.NacosConfig{}
	Trans                        ut.Translator
	RecruitCounterServiceServCon proto.RecruitCounterServiceClient
	UserPostServCon              proto.UserPostClient
	CompanyServCon               proto.CompanyClient
	DepartmentServCon            proto.DepartmentClient
	PostServCon                  proto.PostClient
	Debug                        bool
)
