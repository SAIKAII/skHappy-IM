package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func Init(path string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}
