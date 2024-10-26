package main

import (
	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/fabioods/go-expert-wheater-lab/internal/handler"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client/viacepclient"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client/weatherclient"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/webserver"
	"github.com/fabioods/go-expert-wheater-lab/internal/usecase"
)

func main() {
	c := configs.LoadConfig(".")
	ws := webserver.NewWebServer(c.Port)
	cepClient := viacepclient.NewViaCepClient(c)
	weatherClient := weatherclient.NewWeatherClient(c)
	useCase := usecase.NewWeatherByCepUseCase(weatherClient, cepClient)
	weatherHandler := handler.NewWeatherByCepHandler(useCase)
	ws.AddHandler("/weather/cep/{cep}", weatherHandler.Handle)
	ws.Start()
}
