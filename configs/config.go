package configs

import "github.com/spf13/viper"

type Config struct {
	WeatherApiURL     string `mapstructure:"WEATHER_API_URL"`
	WeatherApiTimeout int    `mapstructure:"WEATHER_API_TIMEOUT"`
	WeatherApiKey     string `mapstructure:"WEATHER_API_KEY"`
	CepApiURL         string `mapstructure:"CEP_API_URL"`
	CepApiTimeout     int    `mapstructure:"CEP_API_TIMEOUT"`
	Port              string `mapstructure:"PORT"`
}

func LoadConfig(path string) *Config {
	var cfg *Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
