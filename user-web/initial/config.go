package initial

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
)

var configDir = "."

func InitConfig() {
	v := viper.New()
	v.AutomaticEnv()
	debug := v.GetBool("DEBUG")

	configPrefix := "config"
	configPrefix = configDir + "/" + "config"
	if debug {
		v.SetConfigFile(fmt.Sprintf("%s-debug.yml", configPrefix))
	} else {
		v.SetConfigFile(fmt.Sprintf("%s-production.yml", configPrefix))
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(global.Config)
	zap.S().Infof("加载配置文件 : %s", global.Config)
	if err != nil {
		panic(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Infof("config change : %s", in.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.Config)
		zap.S().Infof("%v\n", global.Config)
	})
}
