package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DATABASEURL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
