package client

import (
	"context"
)

//go:generate mockery --all --case=underscore --disable-version-string
type CepClient interface {
	AddressByCep(ctx context.Context, cep string) (string, error)
}

type WeatherClient interface {
	WeatherByCity(tx context.Context, city string) (float64, error)
}
