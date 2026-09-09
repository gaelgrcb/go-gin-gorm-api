package config

import (
	"CRUD/pkg/database"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	// Server
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`

	//DB
	DB database.DBConfig `mapstructure:",squash"`

	// Security
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	JWTExpiresHours int    `mapstructure:"JWT_EXPIRES_HOURS"`

	// CORS for Frontend
	CORSAllowedOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config, %v", err)
	}

	if config.DB.Host == "" || config.DB.Name == "" {
		return nil, fmt.Errorf("error some configs are missing or do not match")
	}

	return &config, nil
}
