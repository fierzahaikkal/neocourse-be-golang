package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl     string
	JWTSecret string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return &Config{
		DBUrl:     viper.GetString("DB_URL"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
