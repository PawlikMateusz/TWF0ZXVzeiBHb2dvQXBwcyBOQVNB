package config

import (
	"github.com/spf13/viper"
)

// env vars
const (
	Port                  = "PORT"
	ApiKey                = "API_KEY"
	MaxConcurrentRequests = "CONCURENT_REQUESTS"
)

// default values
const (
	DefaultPort                  = 8080
	DefaultApiKey                = "DEMO_KEY"
	DefaultMaxConcurrentRequests = 10
)

func SetDefaults() {
	viper.SetDefault(Port, DefaultPort)
	viper.SetDefault(ApiKey, DefaultApiKey)
	viper.SetDefault(MaxConcurrentRequests, DefaultMaxConcurrentRequests)
}

func LoadEnvVars() error {
	envVars := []string{
		Port,
		ApiKey,
		MaxConcurrentRequests,
	}
	for _, env := range envVars {
		err := viper.BindEnv(env)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPort() int {
	return viper.GetInt(Port)
}

func GetApiKey() string {
	return viper.GetString(ApiKey)
}

func GetMaxConcurrentRequests() int {
	return viper.GetInt(MaxConcurrentRequests)
}
