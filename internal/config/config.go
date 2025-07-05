package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host                  string              `mapstructure:"host"`
	Port                  int                 `mapstructure:"port"`
	Name                  string              `mapstructure:"name"`
	Username              string              `mapstructure:"username"`
	Password              string              `mapstructure:"password"`
	Options               map[string][]string `mapstructure:"options"`
	MaxIdleConnection     int                 `mapstructure:"max_idle_connection"`
	MaxOpenConnection     int                 `mapstructure:"max_open_connection"`
	MaxConnectionLifetime int                 `mapstructure:"max_connection_lifetime"`
	BatchSize             int                 `mapstructure:"batch_size"`
	SlowThreshold         int                 `mapstructure:"slow_threshold"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	v.SetDefault("server.port", "8080")

	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	_ = v.BindEnv("server.port")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Warning: Failed to read config file: %v\n", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
