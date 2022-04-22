package config

type RecruitSrvInfo struct {
	Name string `mapstructure:"name" json:"name"`
}

type CompanyServConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JwtConfig struct {
	Key string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Host           string            `mapstructure:"host" json:"host"`
	Port           int               `mapstructure:"port" json:"port"`
	Name           string            `mapstructure:"name" json:"name"`
	Tag            []string          `mapstructure:"tag" json:"tag"`
	RecruitSrvInfo RecruitSrvInfo    `mapstructure:"recruit_srv" json:"recruit_srv"`
	CompanySrvInfo CompanyServConfig `mapstructure:"company_srv" json:"company_srv"`
	Local          string            `mapstructure:"local" json:"local"`
	JwtConfig      JwtConfig         `mapstructure:"jwt_config" json:"jwt_config"`
	ConsulConfig   ConsulConfig      `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"nacos_host"`
	Port      int    `mapstructure:"nacos_port"`
	Namespace string `mapstructure:"nacos_namespace"`
	DataId    string `mapstructure:"nacos_data_id"`
	Group     string `mapstructure:"nacos_group"`
	Username  string `mapstructure:"nacos_name"`
	Password  string `mapstructure:"nacos_password"`
}
