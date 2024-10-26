package viacepclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/fabioods/go-expert-wheater-lab/pkg/trace"
)

type ViaCepClient struct {
	CepApiURL     string
	CepApiTimeout int
}

type ViaCepResponse struct {
	Location string `json:"localidade"`
}

func NewViaCepClient(config *configs.Config) *ViaCepClient {
	return &ViaCepClient{
		CepApiURL:     config.CepApiURL,
		CepApiTimeout: config.CepApiTimeout,
	}
}

func (v *ViaCepClient) AddressByCep(ctx context.Context, cep string) (string, error) {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, time.Duration(v.CepApiTimeout)*time.Millisecond)
	defer cancel()
	path := fmt.Sprintf("%s%s/json", v.CepApiURL, cep)
	req, err := http.NewRequestWithContext(contextWithTimeOut, http.MethodGet, path, nil)
	if err != nil {
		return "", errorformated.UnexpectedError(trace.GetTrace(), "error_creating_request", "error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errorformated.UnexpectedError(trace.GetTrace(), "error_requesting_address", "error requesting address: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return "", errorformated.NotFoundError(trace.GetTrace(), "address_not_found", "address not found for cep: %s", cep)
	}

	bytesResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errorformated.UnexpectedError(trace.GetTrace(), "error_reading_response", "error reading response: %v", err)
	}

	var viaCepResponse ViaCepResponse
	err = json.Unmarshal(bytesResponse, &viaCepResponse)
	if err != nil {
		return "", errorformated.UnexpectedError(trace.GetTrace(), "error_unmarshalling_response_cep", "error unmarshalling response cep: %v", err)
	}

	return viaCepResponse.Location, nil
}
