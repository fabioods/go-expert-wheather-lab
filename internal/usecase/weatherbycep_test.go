package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client/mocks"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWeatherByCepUseCase_Execute_Success(t *testing.T) {
	weatherClientMock := new(mocks.WeatherClient)
	cepClientMock := new(mocks.CepClient)
	useCase := NewWeatherByCepUseCase(weatherClientMock, cepClientMock)

	input := InputDTO{Cep: "12345678"}
	expectedCity := "TestCity"
	expectedTemperature := 25.0

	// Mock para cepClient
	cepClientMock.On("AddressByCep", mock.Anything, "12345678").Return(expectedCity, nil)

	// Mock para weatherClient
	weatherClientMock.On("WeatherByCity", mock.Anything, expectedCity).Return(expectedTemperature, nil)

	output, err := useCase.Execute(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, "12345678", output.Cep)
	assert.Equal(t, expectedTemperature, output.CelsiusTemperature)
	assert.Equal(t, (expectedTemperature*1.8)+32, output.FahrenheitTemperature)
	assert.Equal(t, expectedTemperature+273, output.KelvinTemperature)

	cepClientMock.AssertExpectations(t)
	weatherClientMock.AssertExpectations(t)
}

func TestWeatherByCepUseCase_Execute_InvalidCep(t *testing.T) {
	weatherClientMock := new(mocks.WeatherClient)
	cepClientMock := new(mocks.CepClient)
	useCase := NewWeatherByCepUseCase(weatherClientMock, cepClientMock)

	input := InputDTO{Cep: "invalid"}

	output, err := useCase.Execute(context.Background(), input)

	assert.Error(t, err)
	assert.Equal(t, "invalid zipcode", err.(*errorformated.ErrorFormated).Code)
	assert.Equal(t, OutputDTO{}, output)
}

func TestWeatherByCepUseCase_Execute_CityNotFound(t *testing.T) {
	weatherClientMock := new(mocks.WeatherClient)
	cepClientMock := new(mocks.CepClient)
	useCase := NewWeatherByCepUseCase(weatherClientMock, cepClientMock)

	input := InputDTO{Cep: "12345678"}

	// Mock para cepClient retornar cidade vazia
	cepClientMock.On("AddressByCep", mock.Anything, "12345678").Return("", nil)

	output, err := useCase.Execute(context.Background(), input)

	assert.Error(t, err)
	assert.Equal(t, "city_not_found", err.(*errorformated.ErrorFormated).Code)
	assert.Equal(t, OutputDTO{}, output)

	cepClientMock.AssertExpectations(t)
}

func TestWeatherByCepUseCase_Execute_WeatherClientError(t *testing.T) {
	weatherClientMock := new(mocks.WeatherClient)
	cepClientMock := new(mocks.CepClient)
	useCase := NewWeatherByCepUseCase(weatherClientMock, cepClientMock)

	input := InputDTO{Cep: "12345678"}
	expectedCity := "TestCity"

	// Mock para cepClient
	cepClientMock.On("AddressByCep", mock.Anything, "12345678").Return(expectedCity, nil)

	// Mock para weatherClient retornar erro
	weatherClientMock.On("WeatherByCity", mock.Anything, expectedCity).Return(0.0, errors.New("weather service error"))

	output, err := useCase.Execute(context.Background(), input)

	assert.Error(t, err)
	assert.EqualError(t, err, "weather service error")
	assert.Equal(t, OutputDTO{}, output)

	cepClientMock.AssertExpectations(t)
	weatherClientMock.AssertExpectations(t)
}
