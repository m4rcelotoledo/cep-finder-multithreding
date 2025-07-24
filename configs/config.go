package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("timeout", "1s")

	config := &Config{}

	// Converte a string do timeout para Duration
	timeoutStr := viper.GetString("timeout")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, err
	}

	config.Timeout = timeout
	return config, nil
}
