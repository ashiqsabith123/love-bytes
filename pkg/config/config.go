package config

import "github.com/spf13/viper"

type Ports struct {
	AuthSvcPort string `mapstructure:"auth-svc-port"`
}

type Config struct {
	PORTS Ports `mapstructure:"ports"`
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
