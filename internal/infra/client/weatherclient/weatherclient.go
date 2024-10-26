package weatherclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/fabioods/go-expert-wheater-lab/pkg/trace"
)

type WeatherClient struct {
	WeatherApiURL     string
	WeatherApiTimeout int
	WeatherApiKey     string
}

type current struct {
	CelsiusTemperature float64 `json:"temp_c"`
}

type WeatherResponse struct {
	Current current `json:"current"`
}

func NewWeatherClient(config *configs.Config) *WeatherClient {
	return &WeatherClient{
		WeatherApiURL:     config.WeatherApiURL,
		WeatherApiTimeout: config.WeatherApiTimeout,
		WeatherApiKey:     config.WeatherApiKey,
	}
}

func (w *WeatherClient) WeatherByCity(ctx context.Context, city string) (float64, error) {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, time.Duration(w.WeatherApiTimeout)*time.Millisecond)
	defer cancel()

	path := fmt.Sprintf("%s?key=%s&q=%s", w.WeatherApiURL, w.WeatherApiKey, url.QueryEscape(city))
	req, err := http.NewRequestWithContext(contextWithTimeOut, http.MethodGet, path, nil)
	if err != nil {
		return 0, errorformated.UnexpectedError(trace.GetTrace(), "error_creating_request", "error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, errorformated.UnexpectedError(trace.GetTrace(), "error_requesting_address", "error requesting address: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, errorformated.UnexpectedError(trace.GetTrace(), "error_response_status", "unexpected response status: %v", res.Status)
	}

	bytesResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, errorformated.UnexpectedError(trace.GetTrace(), "error_reading_response", "error reading response: %v", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(bytesResponse, &weatherResponse)
	if err != nil {
		return 0, errorformated.UnexpectedError(trace.GetTrace(), "error_unmarshalling_response_weather", "error unmarshalling response weather: %v", err)
	}

	return weatherResponse.Current.CelsiusTemperature, nil
}
