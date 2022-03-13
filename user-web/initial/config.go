package initial

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"hr-saas-go/user-web/global"
)

var configDir = "."

func InitConfig() {
	//v := viper.New()
	//v.AutomaticEnv()
	//debug := v.GetBool("DEBUG")
	//
	//configPrefix := "config"
	//configPrefix = configDir + "/" + "config"
	//if debug {
	//	v.SetConfigFile(fmt.Sprintf("%s-debug.yml", configPrefix))
	//} else {
	//	v.SetConfigFile(fmt.Sprintf("%s-production.yml", configPrefix))
	//}
	//err := v.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}
	//err = v.Unmarshal(global.Config)
	//zap.S().Infof("加载配置文件 : %s", global.Config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	zap.S().Infof("config change : %s", in.Name)
	//	_ = v.ReadInConfig()
	//	_ = v.Unmarshal(global.Config)
	//	zap.S().Infof("%v\n", global.Config)
	//})

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
	// 读取nacos文件
	err = v.Unmarshal(global.NacosConfig)
	zap.S().Infof("加载nacos配置文件 : %s", global.NacosConfig)
	if err != nil {
		panic(err)
	}

	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         "92eb12b0-1a27-41ec-a6f8-4a8bb1611e56", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})
	fmt.Println("GetConfig,config :" + content)

	// json
	err = json.Unmarshal([]byte(content), &global.Config)
}
