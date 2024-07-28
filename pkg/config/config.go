package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	BrokerUrl string
	UserDatabaseUrl string
	WalletDatabaseUrl string
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}