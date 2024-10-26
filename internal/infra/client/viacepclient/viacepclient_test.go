package viacepclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func setup() *ViaCepClient {
	config := &configs.Config{
		CepApiURL:     "http://example.com/",
		CepApiTimeout: 1000,
	}
	return NewViaCepClient(config)
}

// Teste de sucesso: resposta válida da API ViaCEP
func TestAddressByCep_Success(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/12345678/json",
		httpmock.NewStringResponder(http.StatusOK, `{"localidade": "TestCity"}`))

	location, err := client.AddressByCep(context.Background(), "12345678")
	assert.NoError(t, err)
	assert.Equal(t, "TestCity", location)
}

// Teste: CEP não encontrado (404)
func TestAddressByCep_NotFound(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/12345678/json",
		httpmock.NewStringResponder(http.StatusNotFound, ""))

	location, err := client.AddressByCep(context.Background(), "12345678")
	assert.Error(t, err)
	assert.Equal(t, "address_not_found", err.(*errorformated.ErrorFormated).Code)
	assert.Empty(t, location)
}

// Teste: erro ao criar requisição
func TestAddressByCep_RequestError(t *testing.T) {
	client := setup()

	// URL inválida para simular erro de criação de requisição
	client.CepApiURL = "://invalid-url"
	location, err := client.AddressByCep(context.Background(), "12345678")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating request")
	assert.Empty(t, location)
}

// Teste: erro na leitura da resposta
func TestAddressByCep_ReadResponseError(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/12345678/json",
		httpmock.NewBytesResponder(http.StatusOK, []byte{}))

	location, err := client.AddressByCep(context.Background(), "12345678")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshalling response cep: unexpected end of JSON input")
	assert.Empty(t, location)
}

// Teste: erro ao fazer o unmarshal da resposta
func TestAddressByCep_UnmarshalError(t *testing.T) {
	client := setup()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// JSON inválido para simular erro de unmarshal
	httpmock.RegisterResponder("GET", "http://example.com/12345678/json",
		httpmock.NewStringResponder(http.StatusOK, `{"invalid_json": }`))

	location, err := client.AddressByCep(context.Background(), "12345678")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshalling response cep")
	assert.Empty(t, location)
}
