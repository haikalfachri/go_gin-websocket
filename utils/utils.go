package utils

import (
	"log"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}
