package handler

import (
	"encoding/json"
	"github.com/fabioods/go-expert-wheater-lab/internal/usecase"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:generate mockery --all --case=underscore --disable-version-string
type WeatherByCepHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type weatherByCepHandler struct {
	useCase usecase.WeatherByCepUseCase
}

func NewWeatherByCepHandler(useCase usecase.WeatherByCepUseCase) *weatherByCepHandler {
	return &weatherByCepHandler{
		useCase: useCase,
	}
}

func (h *weatherByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	input := usecase.InputDTO{}
	input.Cep = chi.URLParam(r, "cep")

	output, err := h.useCase.Execute(r.Context(), input)
	if err != nil {
		statusCode := err.(*errorformated.ErrorFormated).StatusCode()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
