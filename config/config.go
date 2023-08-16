package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func SetPath(path string) {
	viper.AddConfigPath(path)
}
func SetName(name string) {
	viper.SetConfigName(name)
}
func SetType(t string) {
	viper.SetConfigType(t)
}

func Init(path, name, t string) {
	var err error
	viper.SetConfigName(name) // name of config file (without extension)
	viper.SetConfigType(t)
	viper.AddConfigPath(path)  // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
}

func init() {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err = viper.ReadInConfig()      // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
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
