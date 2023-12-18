package config

import "github.com/spf13/viper"

type Ports struct {
	AuthSvcPort  string `mapstructure:"auth-svc-port"`
	MatchSvcPort string `mapstructure:"match-svc-port"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secret-key"`
}

type Config struct {
	PORTS Ports     `mapstructure:"ports"`
	JWT   JWTConfig `mapstructure:"jwt"`
}

var config Config

func LoadConfig() (Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("pkg/config/")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetSecretKey() string {
	return config.JWT.SecretKey
}
