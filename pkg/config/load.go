package config

import (
	"github.com/spf13/viper"
)

func loadFile() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// LoadRun load config from yml
func LoadRun() error {
	return loadFile()
}
