package config

import (
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	viper.ReadInConfig()            // Find and read the config file
}

func GetString(s string) string {
	if value := viper.GetString(s); value != "" {
		return value
	}
	if value := os.Getenv(s); value != "" {
		return value
	}
	return ""
}
