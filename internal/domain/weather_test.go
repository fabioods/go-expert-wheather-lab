package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefineCep(t *testing.T) {
	weather := NewWeather()

	err := weather.DefineCep("12345678")
	assert.NoError(t, err, "CEP válido deve ser aceito")
	assert.Equal(t, "12345678", weather.Cep, "O CEP deve ser armazenado corretamente")

	err = weather.DefineCep("")
	assert.Error(t, err, "CEP vazio deve retornar erro")
	assert.Equal(t, "cep is required", err.Error(), "Mensagem de erro deve ser 'cep is required'")

	err = weather.DefineCep("12345")
	assert.Error(t, err, "CEP inválido deve retornar erro")
	assert.Equal(t, "cep is invalid", err.Error(), "Mensagem de erro deve ser 'cep is invalid'")
}

func TestDefineCelsiusTemperature(t *testing.T) {
	weather := NewWeather()
	weather.DefineCelsiusTemperature(25.5)
	assert.Equal(t, 25.5, weather.CelsiusTemperature)
}

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	weather := NewWeather()
	weather.DefineCelsiusTemperature(25)
	weather.ConvertCelsiusToFahrenheit()
	assert.Equal(t, 77.0, weather.FahrenheitTemperature)
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	weather := NewWeather()
	weather.DefineCelsiusTemperature(25)
	weather.ConvertCelsiusToKelvin()
	assert.Equal(t, 298.0, weather.KelvinTemperature)
}
