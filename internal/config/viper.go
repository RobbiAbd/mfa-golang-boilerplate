package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigType("env")

	err := config.ReadInConfig()
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	return config
}
