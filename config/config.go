package config

import (
	"os"

	"github.com/spf13/viper"
)

func init() {
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

func GetStringWithConvention(s string, def string) string {
	if value := GetString(s); value != "" {
		return value
	}
	return def
}

func Convention[T any](value *T, def T) T {
	if value == nil {
		return def
	}
	return *value
}
