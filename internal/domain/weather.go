package domain

import (
	"errors"
	"regexp"
)

type Weather struct {
	Cep                   string  `json:"cep"`
	CelsiusTemperature    float64 `json:"celsius_temperature"`
	FahrenheitTemperature float64 `json:"fahrenheit_temperature"`
	KelvinTemperature     float64 `json:"kelvin_temperature"`
}

func NewWeather() *Weather {
	return &Weather{}
}

func (w *Weather) DefineCep(cep string) error {
	err := w.ValidateCep(cep)
	if err != nil {
		return err
	}
	w.Cep = cep
	return nil
}

func (w *Weather) ValidateCep(cep string) error {
	if cep == "" {
		return errors.New("cep is required")
	}
	cepRegex := `^\d{8}$`
	matched, err := regexp.MatchString(cepRegex, cep)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("cep is invalid")
	}
	return nil
}

func (w *Weather) DefineCelsiusTemperature(celsiusTemperature float64) {
	w.CelsiusTemperature = celsiusTemperature
}

func (w *Weather) ConvertCelsiusToFahrenheit() {
	w.FahrenheitTemperature = (w.CelsiusTemperature * 1.8) + 32
}

func (w *Weather) ConvertCelsiusToKelvin() {
	w.KelvinTemperature = w.CelsiusTemperature + 273
}
