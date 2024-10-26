package weatherclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func setup() *WeatherClient {
	config := &configs.Config{
		WeatherApiURL:     "http://example.com",
		WeatherApiTimeout: 1000,
		WeatherApiKey:     "test-key",
	}
	return NewWeatherClient(config)
}

// Teste de sucesso: resposta válida da API de clima
func TestWeatherByCity_Success(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com?key=test-key&q=TestCity",
		httpmock.NewStringResponder(http.StatusOK, `{"current": {"temp_c": 25.0}}`))

	temp, err := client.WeatherByCity(context.Background(), "TestCity")
	assert.NoError(t, err)
	assert.Equal(t, 25.0, temp)
}

// Teste: erro ao criar requisição
func TestWeatherByCity_RequestError(t *testing.T) {
	client := setup()

	// Configura uma URL inválida para simular erro de criação de requisição
	client.WeatherApiURL = "://invalid-url"

	temp, err := client.WeatherByCity(context.Background(), "TestCity")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating request")
	assert.Equal(t, 0.0, temp)
}

// Teste: erro de status na resposta
func TestWeatherByCity_ResponseStatusError(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com?key=test-key&q=TestCity",
		httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	temp, err := client.WeatherByCity(context.Background(), "TestCity")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected response status")
	assert.Equal(t, 0.0, temp)
}

// Teste: erro de leitura da resposta
func TestWeatherByCity_ReadResponseError(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Retorna uma resposta com corpo vazio para simular erro de leitura
	httpmock.RegisterResponder("GET", "http://example.com?key=test-key&q=TestCity",
		httpmock.NewBytesResponder(http.StatusOK, []byte{}))

	temp, err := client.WeatherByCity(context.Background(), "TestCity")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshalling response weather: unexpected end of JSON input")
	assert.Equal(t, 0.0, temp)
}

// Teste: erro de unmarshal da resposta JSON
func TestWeatherByCity_UnmarshalError(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// JSON inválido para simular erro de unmarshal
	httpmock.RegisterResponder("GET", "http://example.com?key=test-key&q=TestCity",
		httpmock.NewStringResponder(http.StatusOK, `{"invalid_json": }`))

	temp, err := client.WeatherByCity(context.Background(), "TestCity")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshalling response weather")
	assert.Equal(t, 0.0, temp)
}
