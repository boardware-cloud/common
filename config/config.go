package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
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

func Convention[T any](value *T, def T) T {
	if value == nil {
		return def
	}
	return *value
}
