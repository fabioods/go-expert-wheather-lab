package errorformated

import (
	"fmt"
	"net/http"

	"github.com/fabioods/go-expert-wheater-lab/pkg/trace"
)

const emptyStatusCode = 0

type ErrorFormated struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Cause   error           `json:"-"`
	Trace   trace.TraceInfo `json:"-"`
}

func (e ErrorFormated) Error() string {
	return e.Message
}

func (e ErrorFormated) StatusCode() int {
	if e.Status == emptyStatusCode {
		return http.StatusInternalServerError
	}
	return e.Status
}

func BadRequestError(trace trace.TraceInfo, code string, message string, msgValues ...interface{}) *ErrorFormated {
	return &ErrorFormated{
		Code:    code,
		Message: fmt.Sprintf(message, msgValues...),
		Status:  http.StatusBadRequest,
		Trace:   trace,
	}
}

func NotFoundError(trace trace.TraceInfo, code string, message string, msgValues ...interface{}) *ErrorFormated {
	return &ErrorFormated{
		Code:    code,
		Message: fmt.Sprintf(message, msgValues...),
		Status:  http.StatusNotFound,
		Trace:   trace,
	}
}

func UnexpectedError(trace trace.TraceInfo, code string, message string, msgValues ...interface{}) *ErrorFormated {
	return &ErrorFormated{
		Code:    code,
		Message: fmt.Sprintf(message, msgValues...),
		Status:  http.StatusInternalServerError,
		Trace:   trace,
	}
}

func UnprocesableEntityError(trace trace.TraceInfo, code string, message string, msgValues ...interface{}) *ErrorFormated {
	return &ErrorFormated{
		Code:    code,
		Message: fmt.Sprintf(message, msgValues...),
		Status:  http.StatusUnprocessableEntity,
		Trace:   trace,
	}
}
