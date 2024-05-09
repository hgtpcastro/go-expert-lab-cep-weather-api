package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	CepApiUrl     string `mapstructure:"CEP_API_URL"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
}

func LoadConfig(path string) (config *Config, err error) {
	var c *Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return &Config{}, err
	}

	return c, nil
}
