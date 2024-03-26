package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token        string
	DatabaseURL  string
	DebugLogging bool
}

var current *Config

func Load() *Config {
	if current != nil {
		return current
	}

	current = &Config{}
	viper.Unmarshal(current)
	return current
}
