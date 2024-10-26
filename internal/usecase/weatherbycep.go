package usecase

import (
	"context"

	"github.com/fabioods/go-expert-wheater-lab/internal/domain"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/fabioods/go-expert-wheater-lab/pkg/trace"
)

type InputDTO struct {
	Cep string `json:"cep"`
}

type OutputDTO struct {
	Cep                   string  `json:"cep"`
	CelsiusTemperature    float64 `json:"celsius_temperature"`
	FahrenheitTemperature float64 `json:"fahrenheit_temperature"`
	KelvinTemperature     float64 `json:"kelvin_temperature"`
}

type WeatherByCepUseCase interface {
	Execute(context context.Context, input InputDTO) (OutputDTO, error)
}

type weatherByCepUseCase struct {
	weatherClient client.WeatherClient
	cepClient     client.CepClient
}

func NewWeatherByCepUseCase(weatherClient client.WeatherClient, cepClient client.CepClient) *weatherByCepUseCase {
	return &weatherByCepUseCase{
		weatherClient: weatherClient,
		cepClient:     cepClient,
	}
}

func (w *weatherByCepUseCase) Execute(context context.Context, input InputDTO) (OutputDTO, error) {
	weather := domain.NewWeather()
	err := weather.DefineCep(input.Cep)
	if err != nil {
		return OutputDTO{}, errorformated.UnprocesableEntityError(trace.GetTrace(), "invalid zipcode", err.Error())
	}

	city, err := w.cepClient.AddressByCep(context, input.Cep)
	if err != nil {
		return OutputDTO{}, err
	}

	if city == "" {
		return OutputDTO{}, errorformated.NotFoundError(trace.GetTrace(), "city_not_found", "can not find zipcode: %s", input.Cep)
	}

	celsiusTemperature, err := w.weatherClient.WeatherByCity(context, city)
	if err != nil {
		return OutputDTO{}, err
	}

	weather.DefineCelsiusTemperature(celsiusTemperature)
	weather.ConvertCelsiusToFahrenheit()
	weather.ConvertCelsiusToKelvin()
	return OutputDTO{
		Cep:                   weather.Cep,
		CelsiusTemperature:    weather.CelsiusTemperature,
		FahrenheitTemperature: weather.FahrenheitTemperature,
		KelvinTemperature:     weather.KelvinTemperature,
	}, nil
}
