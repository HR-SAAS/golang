package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JwtConfig struct {
	Key string `mapstructure:"key"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
	Local       string        `mapstructure:"local"`
	JwtConfig   JwtConfig     `mapstructure:"jwt_config"`
}
