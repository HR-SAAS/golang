package config

type UserSrvConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JwtConfig struct {
	Key string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	Name         string        `mapstructure:"name"`
	UserSrvInfo  UserSrvConfig `mapstructure:"user_srv"`
	Local        string        `mapstructure:"local"`
	JwtConfig    JwtConfig     `mapstructure:"jwt_config"`
	ConsulConfig ConsulConfig  `mapstructure:"consul"`
}
