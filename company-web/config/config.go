package config

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type CompanySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JwtConfig struct {
	Key string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Host           string           `mapstructure:"host" json:"host"`
	Port           int              `mapstructure:"port" json:"port"`
	Name           string           `mapstructure:"name" json:"name"`
	UserSrvInfo    UserSrvConfig    `mapstructure:"user_srv" json:"user_srv"`
	CompanySrvInfo CompanySrvConfig `mapstructure:"company_srv" json:"company_srv"`
	Local          string           `mapstructure:"local" json:"local"`
	JwtConfig      JwtConfig        `mapstructure:"jwt_config" json:"jwt_config"`
	ConsulConfig   ConsulConfig     `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"data-id"`
	Group     string `mapstructure:"group"`
}
