package config

type UserSrvConfig struct {
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

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Host         string        `mapstructure:"host" json:"host"`
	Port         int           `mapstructure:"port" json:"port"`
	Name         string        `mapstructure:"name" json:"name"`
	UserSrvInfo  UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	Local        string        `mapstructure:"local" json:"local"`
	JwtConfig    JwtConfig     `mapstructure:"jwt_config" json:"jwt_config"`
	ConsulConfig ConsulConfig  `mapstructure:"consul" json:"consul"`
	RedisConfig  RedisConfig   `mapstructure:"redis" json:"redis"`
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
