package handler

import (
	"context"
	"encoding/json"
	_ "errors"
	"github.com/fabioods/go-expert-wheater-lab/pkg/trace"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabioods/go-expert-wheater-lab/internal/usecase"
	"github.com/fabioods/go-expert-wheater-lab/internal/usecase/mocks"
	"github.com/fabioods/go-expert-wheater-lab/pkg/errorformated"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWeatherByCepHandler_Success(t *testing.T) {
	mockUseCase := new(mocks.WeatherByCepUseCase)
	handler := NewWeatherByCepHandler(mockUseCase)

	input := usecase.InputDTO{Cep: "12345678"}
	expectedOutput := usecase.OutputDTO{CelsiusTemperature: 25.0}
	mockUseCase.On("Execute", mock.Anything, input).Return(expectedOutput, nil)

	req := httptest.NewRequest("GET", "/weather/12345678", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "12345678")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var output usecase.OutputDTO
	json.NewDecoder(w.Body).Decode(&output)
	assert.Equal(t, expectedOutput, output)
	mockUseCase.AssertExpectations(t)
}

func TestWeatherByCepHandler_Failure(t *testing.T) {
	mockUseCase := new(mocks.WeatherByCepUseCase)
	handler := NewWeatherByCepHandler(mockUseCase)

	input := usecase.InputDTO{Cep: "invalid-cep"}
	expectedError := errorformated.BadRequestError(trace.GetTrace(), "invalid zipcode", "cep is invalid")
	mockUseCase.On("Execute", mock.Anything, input).Return(usecase.OutputDTO{}, expectedError)

	req := httptest.NewRequest("GET", "/weather/invalid-cep", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "invalid-cep")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.Handle(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errResp errorformated.ErrorFormated
	json.NewDecoder(w.Body).Decode(&errResp)
	assert.Equal(t, expectedError.Error(), errResp.Error())
	mockUseCase.AssertExpectations(t)
}
