package configs

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg := LoadConfig("mock")

	assert.Equal(t, "https://viacep.com.br/ws/", cfg.CepApiURL)
	assert.Equal(t, 1000, cfg.CepApiTimeout)
	assert.Equal(t, "https://api.weatherapi.com/v1/current.json", cfg.WeatherApiURL)
	assert.Equal(t, 5000, cfg.WeatherApiTimeout)
	assert.Equal(t, "X", cfg.WeatherApiKey)
	assert.Equal(t, "8080", cfg.Port)
}

func TestLoadConfigWithMissingEnvVars(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	assert.Panics(t, func() {
		LoadConfig("")
	})
}
